package models

import "github.com/jinzhu/gorm"

func init() {
	GetDBInstance().AutoMigrate(&Role{}, &RoleUserRelation{}, &RolePermissionRelation{})
}

type Role struct {
	gorm.Model

	Name string `gorm:"type:varchar(64);index:idx_name;not null;default:'';unique_index:idx_name"`
	Slug string `gorm:"type:varchar(64); not null"`
}

func (r *Role) TableName() string {
	return `role_t`
}

type RoleUserRelation struct {
	gorm.Model

	RoleID uint `gorm:"column:role_id; not null"`
	UserID uint `gorm:"column:user_id; not null"`
}

func (r *RoleUserRelation) TableName() string {
	return `role_user_rel_t`
}

type RolePermissionRelation struct {
	gorm.Model

	RoleID uint `gorm:"column:role_id; not null"`
	PermID uint `gorm:"column:perm_id; not null"`
}

func (r *RolePermissionRelation) TableName() string {
	return `role_permission_rel_t`
}

// TODO: table name
func GetPermissionsByUserId(userId uint) ([]Permission, error) {
	permissions, roleIds := make([]Permission, 0), make([]uint, 0)

	err := GetDBInstance().Table(`role_user_rel_t`).Select(`user_id=?`, userId).Find(&roleIds).Error
	if err != nil {
		return permissions, err
	}

	err = GetDBInstance().Table(`role_permission_rel_t`).Where(`role_id in (?)`, roleIds).Find(&permissions).Error
	if err != nil {
		return permissions, err
	}

	return permissions, nil
}
