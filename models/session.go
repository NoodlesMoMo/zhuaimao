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

func (helper *SessionHelper) GetUserId(cookieSec string) (id string) {
	var ok bool
	if id, ok = helper.ses.Get(cookieSec).(string); ok {
		return
	}

	return ""
}

func (helper *SessionHelper) Set(value interface{}) string {
	key := GenerateSessionId()
	helper.ses.Set(key, value)

	return key
}

func GenerateSessionId() string {
	return uuid.New().String()
}
