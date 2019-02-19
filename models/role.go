package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	Model
	Name string `json:"name"`
	Menu []Menu `json:"menu" gorm:"many2many:role_menu;"`
}

func ExistRoleByID(id int) (bool, error) {
	var role Role
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if role.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetRoleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Role{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetRoles(pageNum int, pageSize int, maps interface{}) ([]*Role, error) {
	var role []*Role
	err := db.Preload("Menu").Where(maps).Offset(pageNum).Limit(pageSize).Find(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return role, nil
}

func GetRole(id int) (*Role, error) {
	var role Role
	err := db.Preload("Menu").Where("id = ? AND deleted_on = ? ", id, 0).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &role, nil
}
func CheckRoleName(name string) (bool, error) {
	var role Role
	err := db.Where("name = ? AND deleted_on = ? ", name, 0).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CheckRoleNameId(name string, id int) (bool, error) {
	var role Role
	err := db.Where("name = ? AND id != ? AND deleted_on = ? ", name, id, 0).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}

	return false, nil
}

func EditRole(id int, data map[string]interface{}) error {
	var role []Role
	var menu Menu
	db.Where("id in (?)", data["menu_id"].(int)).Find(&menu)

	if err := db.Where("id = ? AND deleted_on = ? ", id, 0).Find(&role).Error; err != nil {
		return err
	}
	db.Model(&role).Association("Menu").Replace(menu)
	db.Model(&role).Update(data)

	return nil
}

func AddRole(data map[string]interface{}) (id int, err error) {
	role := Role{
		Name: data["name"].(string),
	}
	var menu []Menu
	db.Where("id in (?)", data["menu_id"].(int)).Find(&menu)
	if err := db.Create(&role).Association("Menu").Append(menu).Error; err != nil {
		return 0, err
	}
	return role.ID, nil
}

func DeleteRole(id int) error {
	var role Role
	db.Where("id = ?", id).Find(&role)
	db.Model(&role).Association("Menu").Delete()
	if err := db.Where("id = ?", id).Delete(&role).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllRole() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Role{}).Error; err != nil {
		return err
	}

	return nil
}

func GetRolesAll() ([]*Role, error) {
	var role []*Role
	err := db.Preload("Menu").Find(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return role, nil
}
