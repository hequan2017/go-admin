package models

import "github.com/jinzhu/gorm"

type Menu struct {
	Model
	Path         string `json:"path"`
	Method         string `json:"method"`
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
		Path:         data["path"].(string),
		Method:         data["method"].(string),
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
