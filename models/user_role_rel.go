package models

import "github.com/jinzhu/gorm"

func init() {
	GetDBInstance().AutoMigrate(&UseRoleRelation{})
}

type UseRoleRelation struct {
	gorm.Model

	UserID uint `gorm:"type:int32;column:user_id; not null"`
	RoleID uint `gorm:"type:int32;column:role_id; not null"`
}

func (r *UseRoleRelation) TableName() string {
	return `user_role_t`
}
