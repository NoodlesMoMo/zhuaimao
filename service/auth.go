package service

import "zhuaimao/models"

func CheckUser(username, password []byte) (models.User, bool) {
	user, ok := models.User{}, false

	err := models.GetDBInstance().Table(`user_t`).First(&user, "name=?", username).Error
	if err != nil {
		return user, false
	} else {
		if user.Password == string(password) {
			ok = true
		}
	}

	return user, ok
}
