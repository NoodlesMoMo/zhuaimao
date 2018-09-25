package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"zhuaimao/models"
)

var (
	defaultWaitDuration = time.Duration(5 * time.Second)

	atExitHooks  = make([]func() error, 0)
	exitHookLock = sync.Mutex{}
)

func init() {
	RegisteExitHook(func() error {
		models.DestroyDBInstance()
		return nil
	})
}

type ListenerWrap struct {
	listener        net.Listener
	maxWaitDuration time.Duration
	stop            chan struct{}
	connCount       uint64
	shutdown        uint64
}

func (ln *ListenerWrap) Accept() (net.Conn, error) {
	conn, err := ln.listener.Accept()
	if err != nil {
		return nil, err
	}

	atomic.AddUint64(&ln.connCount, 1)

	return &connWrap{Conn: conn, listener: ln}, nil
}

func (ln *ListenerWrap) Addr() net.Addr {
	return ln.listener.Addr()
}

func (ln *ListenerWrap) Close() error {
	if err := ln.listener.Close(); err != nil {
		return err
	}

	return ln.waitForZeroConns()
}

func (ln *ListenerWrap) waitForZeroConns() error {
	atomic.AddUint64(&ln.shutdown, 1)

	if atomic.LoadUint64(&ln.connCount) == 0 {
		close(ln.stop)
		return nil
	}

	select {
	case <-ln.stop:
		return nil
	case <-time.After(ln.maxWaitDuration):
		return errors.New("close all stablished connection failed")
	}

	return nil
}

func (ln *ListenerWrap) tryCloseConnection() {
	cnt := atomic.AddUint64(&ln.connCount, ^uint64(0))

	if atomic.LoadUint64(&ln.shutdown) != 0 && cnt == 0 {
		close(ln.stop)
	}
}

type connWrap struct {
	net.Conn
	listener *ListenerWrap
}

func (c *connWrap) Close() error {
	err := c.Conn.Close()
	if err != nil {
		return err
	}

	c.listener.tryCloseConnection()

	return nil
}

func GracefulListenEx(address string, maxWait time.Duration) (net.Listener, error) {
	ln, err := net.Listen("tcp4", address)
	if err != nil {
		return nil, err
	}

	return &ListenerWrap{
		listener:        ln,
		stop:            make(chan struct{}),
		maxWaitDuration: maxWait,
	}, nil
}

func GracefulListen(address string) (net.Listener, error) {
	return GracefulListenEx(address, defaultWaitDuration)
}

func RegisteExitHook(fn func() error) {
	exitHookLock.Lock()
	atExitHooks = append(atExitHooks, fn)
	exitHookLock.Unlock()
}

func signalAction(ln net.Listener) {
	sc := make(chan os.Signal)
	signal.Notify(sc)

	for s := range sc {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			code := 0
			for _, fn := range atExitHooks {
				if e := fn(); e != nil {
					code += 1
				}
			}

			if e := ln.Close(); e != nil {
				code += 1
				fmt.Fprintf(os.Stderr, "close listener error: %s\n", e.Error())
			}

			fmt.Println(">>>>> exit code:", code)
			os.Exit(code)
		}
	}
}
