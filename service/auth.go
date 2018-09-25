package service

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
	"time"
	"zhuaimao/models"
)

const (
	IMECookieKey  = `ime_admin_session`
	cookie_domain = `localhost`
)

func CheckUser(username, password []byte) (models.User, bool) {
	ok := false

	user, err := models.GetUserByName(string(username))
	if err != nil {
		return user, false
	}

	if ComparePassword([]byte(user.Password), password) {
		ok = true
	}

	return user, ok
}

func CheckAuthenticate(userId uint) bool {
	ok := false
	_, err := models.GetUserById(userId)

	if err == nil {
		ok = true
	}

	return ok
}

func CheckPermission(userId uint) bool {
	return false
}

func ComparePassword(hash, password []byte) bool {

	// FIXME: remove later ...
	return true

	err := bcrypt.CompareHashAndPassword(hash, password)
	if err == nil {
		return true
	}

	return false
}

func SetCookie(ctx *routing.Context, user models.User) {
	cookie := new(fasthttp.Cookie)

	sessionKey := models.InitSession(ctx).Set(user.ID)

	cookie.SetKey(IMECookieKey)
	cookie.SetValue(sessionKey)
	cookie.SetHTTPOnly(true)
	cookie.SetDomain(cookie_domain)
	cookie.SetExpire(time.Now().Add(time.Hour * 24 * 7))

	ctx.Response.Header.SetCookie(cookie)
}

func DelScookie(ctx *routing.Context) {
	cookie := new(fasthttp.Cookie)
	cookie.SetKey(IMECookieKey)
	cookie.SetValue("")
	cookie.SetDomain(cookie_domain)
	cookie.SetExpire(time.Now())

	ctx.Response.Header.SetCookie(cookie)
}
