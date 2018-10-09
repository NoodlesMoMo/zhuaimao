package models

import "github.com/jinzhu/gorm"

func init() {
	GetDBInstance().AutoMigrate(&Menu{})
}

type Menu struct {
	gorm.Model

	ParentID uint   `gorm:"not null; column:parent_id"`
	Title    string `gorm:"not null"`
	Icon     string `gorm:"not null"`
	URI      string `gorm:"not null; default:'/';"`
	RoleIds  string `gorm:"not null; column:role_ids"`
}

func (p *Menu) TableName() string {
	return `menu_t`
}
