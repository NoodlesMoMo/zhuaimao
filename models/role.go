package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func init() {
	GetDBInstance().AutoMigrate(&Role{}, &RoleGroupRelation{}, &RolePermissionRelation{})
}

type Role struct {
	gorm.Model

	Name string `gorm:"type:varchar(64);index:idx_name;not null;default:'';unique_index:idx_name"`
}

func (r *Role) TableName() string {
	return `role_t`
}

func (r *Role) Add(name string) (Role, error) {
	result := Role{}

	cnt := 0
	GetDBInstance().Table(r.TableName()).Where(`deleted_at is null`).Where(`name`, name).Count(&cnt)
	if cnt > 0 {
		return result, fmt.Errorf("%s has exist", name)
	}

	err := GetDBInstance().Table(r.TableName()).Create(&Role{Name: name}).Error
	if err != nil {
		return result, err
	}

	err = GetDBInstance().Where("name=?", name).First(&result).Error

	return result, err
}

func (r *Role) List(page, psize int) ([]Role, error) {
	result := make([]Role, 0)

	if page > 0 {
		page -= 1
	}

	if psize <= 0 {
		psize = 10
	}

	err := GetDBInstance().Limit(psize).Offset(page * psize).Find(&result).Error

	return result, err
}

func (r *Role) All() ([]Role, error) {
	result := make([]Role, 0)

	err := GetDBInstance().Table(r.TableName()).Where("deleted_at is null").Find(&result).Error

	return result, err
}

type RoleGroupRelation struct {
	gorm.Model

	RoleID  uint `gorm:"column:role_id; not null"`
	GroupID uint `gorm:"column:group_id; not null"`
}

func (r *RoleGroupRelation) TableName() string {
	return `role_group_rel_t`
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
