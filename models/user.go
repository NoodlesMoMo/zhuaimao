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

	Name      string `gorm:"type:varchar(64);index:idx_name;not null;default:''"`
	Password  string `gorm:"type:varchar(84);not null;default:''"`
	Level     string `gorm:"type:varchar(64); not null; default:''"`
	LevelName string `gorm:"type:varchar(64); not null; default:''"`
	Avatar    string `gorm:"type:varchar(255)"`
	Token     string `gorm:"size:64"`

	Permissions []Permission `gorm:"-"`
}

type Permission struct {
	Method []string
	Path   []string
}

func (u *User) TableName() string {
	return `user_t`
}
