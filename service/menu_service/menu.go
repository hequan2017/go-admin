package menu_service

import (
	"github.com/casbin/casbin"
	"github.com/hequan2017/go-admin/models"
)

type Menu struct {
	ID     int
	Path   string
	Method string

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Menu     *models.Menu     `inject:""`
	Enforcer *casbin.Enforcer `inject:""`
}

func (a *Menu) Add() error {
	menu := map[string]interface{}{
		"path":   a.Path,
		"method": a.Method,
	}
	if err := models.AddMenu(menu); err != nil {
		return err
	}

	return nil
}

func (a *Menu) Edit() error {
	return models.EditMenu(a.ID, map[string]interface{}{
		"path":   a.Path,
		"method": a.Method,
	})
}

func (a *Menu) Get() (*models.Menu, error) {

	menu, err := models.GetMenu(a.ID)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (a *Menu) GetAll() ([]*models.Menu, error) {
	Menu, err := models.GetMenus(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return Menu, nil
}

func (a *Menu) Delete() error {
	return models.DeleteMenu(a.ID)
}

func (a *Menu) ExistByID() (bool, error) {
	return models.ExistMenuByID(a.ID)
}

func (a *Menu) Count() (int, error) {
	return models.GetMenuTotal(a.getMaps())
}

func (a *Menu) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}
