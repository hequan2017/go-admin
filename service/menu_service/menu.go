package menu_service

import (
	"github.com/hequan2017/go-admin/models"
)

type Menus struct {
	ID   int
	Path string
	Method string

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int
}

func (a *Menus) Add() error {
	menu := map[string]interface{}{
		"path":       a.Path,
		"created_by": a.CreatedBy,
	}

	if err := models.AddMenu(menu); err != nil {
		return err
	}

	return nil
}

func (a *Menus) Edit() error {
	return models.EditMenu(a.ID, map[string]interface{}{
		"path":       a.Path,
		"created_by":  a.CreatedBy,
		"modified_by": a.ModifiedBy,
	})
}

func (a *Menus) Get() (*models.Menu, error) {

	menu, err := models.GetMenu(a.ID)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (a *Menus) GetAll() ([]*models.Menu, error) {
	menus, err := models.GetMenus(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (a *Menus) Delete() error {
	return models.DeleteMenu(a.ID)
}

func (a *Menus) ExistByID() (bool, error) {
	return models.ExistMenuByID(a.ID)
}

func (a *Menus) Count() (int, error) {
	return models.GetMenuTotal(a.getMaps())
}

func (a *Menus) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}
