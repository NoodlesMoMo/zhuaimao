package models

import "github.com/jinzhu/gorm"

func init() {
	GetDBInstance().AutoMigrate(&Role{})
}

type Role struct {
	gorm.Model

	Name string `gorm:"type:varchar(64);index:idx_name;not null;default:'';unique_index:idx_name"`
	Slug string `gorm:"type:varchar(64); not null"`
}

func (r *Role) TableName() string {
	return `role_t`
}
