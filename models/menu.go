package models

import (
	"github.com/jinzhu/gorm"
)

type Menu struct {
	Model
	Name   string `json:"name"`
	Type   string `json:"type"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func ExistMenuByID(id int) (bool, error) {
	var menu Menu
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if menu.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetMenuTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Menu{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetMenus(pageNum int, pageSize int, maps interface{}) ([]*Menu, error) {
	var menu []*Menu
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return menu, nil
}

func GetMenu(id int) (*Menu, error) {
	var menu Menu
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&menu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &menu, nil
}

func EditMenu(id int, data interface{}) error {
	if err := db.Model(&Menu{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddMenu(data map[string]interface{}) error {
	menu := Menu{
		Name:   data["name"].(string),
		Path:   data["path"].(string),
		Method: data["method"].(string),
	}
	if err := db.Create(&menu).Error; err != nil {
		return err
	}

	return nil
}

func DeleteMenu(id int) error {
	if err := db.Where("id = ?", id).Delete(Menu{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllMenu() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Menu{}).Error; err != nil {
		return err
	}
	return nil
}

func EditMenuGetRoles(id int) []int {
	var menu Menu
	var role []Role

	db.Model(&menu).Where("id = ? AND deleted_on = ? ", id, 0)
	db.Joins(" left join go_role_menu b on go_role.id=b.role_id left join go_menu c on c.id=b.menu_id").Where("c.id = ?", id).Find(&role)

	roleList := []int{}
	for _, v := range role {
		roleList = append(roleList, v.ID)
	}
	return roleList
}
