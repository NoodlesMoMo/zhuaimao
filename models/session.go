package models

import (
	"github.com/fasthttp-contrib/sessions"
	"github.com/google/uuid"
	"github.com/qiangxue/fasthttp-routing"
)

type SessionHelper struct {
	ses sessions.Session
}

func InitSession(ctx *routing.Context) *SessionHelper {
	return &SessionHelper{
		ses: sessions.StartFasthttp(ctx.RequestCtx),
	}
}

func (helper *SessionHelper) GetUserId(cookieSec string) (id uint) {
	var ok bool
	if id, ok = helper.ses.Get(cookieSec).(uint); ok {
		return
	}

	return 0
}

func (helper *SessionHelper) Set(value interface{}) string {
	key := GenerateSessionId()
	helper.ses.Set(key, value)

	return key
}

func GenerateSessionId() string {
	return uuid.New().String()
}
