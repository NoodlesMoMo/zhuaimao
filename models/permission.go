package models

import "github.com/jinzhu/gorm"

func init() {
	GetDBInstance().AutoMigrate(&Permission{})
}

type Permission struct {
	gorm.Model

	Name   string `gorm:"type:varchar(64);index:idx_name;not null;default:'';unique_index:idx_name"`
	Slug   string `gorm:"type:varchar(64);not null;default:''"`
	Method string `gorm:"type:varchar(8);not null"`
	Path   string `gorm:"type:varchar(255);not null"`
}

func (p *Permission) TableName() string {
	return `permission_t`
}
