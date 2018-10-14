package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

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

func (p *Permission) List(page, psize int) ([]Permission, error) {
	result := make([]Permission, 0)

	if page > 0 {
		page -= 1
	}

	if psize <= 0 {
		psize = 10
	}

	err := GetDBInstance().Limit(psize).Offset(page * psize).Find(&result).Error

	return result, err
}

func (p *Permission) Add(name, slug, method, path string) (Permission, error) {
	result := Permission{}

	cnt := 0
	GetDBInstance().Table(p.TableName()).Where(`deleted_at is null`).Where(`name`, name).Count(&cnt)
	if cnt > 0 {
		return result, fmt.Errorf("%s has exist", name)
	}

	err := GetDBInstance().Table(p.TableName()).Create(&Permission{
		Name:   name,
		Slug:   slug,
		Method: method,
		Path:   path,
	}).Error
	if err != nil {
		return result, err
	}

	err = GetDBInstance().Where("name=?", name).First(&result).Error

	return result, err
}

func (p *Permission) Delete(id uint) error {
	return nil
}

func (p *Permission) Edit(perm *Permission) error {
	return nil
}
