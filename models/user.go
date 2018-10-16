package models

import (
	"github.com/jinzhu/gorm"

	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	GetDBInstance().AutoMigrate(&User{}, &Group{})
}

type Group struct {
	gorm.Model

	Name  string `gorm:"type:varchar(64);index:idx_name;not null;default:''"`
	Roles string `gorm:"default:''"`
}

func (g *Group) TableName() string {
	return `group_t`
}

func (g *Group) Add(name, roles string) error {
	result := Group{}

	cnt := 0
	GetDBInstance().Table(g.TableName()).Where(`deleted_at is null`).Where(`name`, name).Count(&cnt)
	if cnt > 0 {
		return fmt.Errorf("%s has exist", name)
	}

	err := GetDBInstance().Table(g.TableName()).Create(&Group{Name: name, Roles: roles}).Error
	if err != nil {
		return err
	}

	err = GetDBInstance().Where("name=?", name).First(&result).Error

	return err
}

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(64);index:idx_name;not null;default:''"`
	Password string `gorm:"type:varchar(84);not null;default:''"`
	Groups   string `gorm:"not null"`
	Avatar   string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255); not null; default:'';unique"`
	Token    string `gorm:"size:64"`
}

func (u *User) TableName() string {
	return `user_t`
}

func (u *User) Add(name, passwd, groups, avatar, email string) error {
	result := User{}

	cnt := 0
	GetDBInstance().Table(u.TableName()).Where(`deleted_at is null`).Where(`name`, name).Count(&cnt)
	if cnt > 0 {
		return fmt.Errorf("%s has exist", name)
	}

	err := GetDBInstance().Table(u.TableName()).Create(&User{
		Name:     name,
		Password: passwd,
		Groups:   groups,
		Avatar:   avatar,
		Email:    email,
	}).Error

	if err != nil {
		return err
	}

	err = GetDBInstance().Where("name=?", name).First(&result).Error

	return err
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
