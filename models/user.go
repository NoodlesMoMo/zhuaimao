package models

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	GetDBInstance().AutoMigrate(&User{})
}

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(64);index:idx_name;not null;default:''"`
	Password string `gorm:"type:varchar(84);not null;default:''"`
	RoleID   uint   `gorm:"not null"`
	Avatar   string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255); not null; default:''"`
	Token    string `gorm:"size:64"`
}

func (u *User) TableName() string {
	return `user_t`
}

func GetUserById(id uint) (User, error) {
	user := User{}

	err := GetDBInstance().First(&user, id).Error

	return user, err
}

func GetUserByName(name string) (User, error) {
	var err error
	user := User{}

	err = GetDBInstance().First(&user).Where(`name=?`, name).Error

	return user, err
}
