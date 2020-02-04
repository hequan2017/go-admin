package menu_service

import (
	"github.com/casbin/casbin"
	"go-admin/models"
	"go-admin/service/role_service"
)

type Menu struct {
	ID     int
	Name   string
	Path   string
	Type   string
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
		"name":   a.Name,
		"type":   a.Type,
		"path":   a.Path,
		"method": a.Method,
	}
	if err := models.AddMenu(menu); err != nil {
		return err
	}

	return nil
}

func (a *Menu) Edit() error {
	err := models.EditMenu(a.ID, map[string]interface{}{
		"name":   a.Name,
		"type":   a.Type,
		"path":   a.Path,
		"method": a.Method,
	})
	if err != nil {
		return err
	}
	roleList := models.EditMenuGetRoles(a.ID)
	roleService := Role_service.Role{}
	for _, v := range roleList {
		err := roleService.LoadPolicy(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Menu) Get() (*models.Menu, error) {

	menu, err := models.GetMenu(a.ID)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (a *Menu) GetAll() ([]*models.Menu, error) {

	if a.ID != 0 {
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		maps["id"] = a.ID
		user, err := models.GetMenus(a.PageNum, a.PageSize, maps)
		if err != nil {
			return nil, err
		}
		return user, nil
	} else {
		Menu, err := models.GetMenus(a.PageNum, a.PageSize, a.getMaps())
		if err != nil {
			return nil, err
		}
		return Menu, nil
	}
}

func (a *Menu) Delete() error {
	err := models.DeleteMenu(a.ID)
	if err != nil {
		return err
	}
	roleList := models.EditMenuGetRoles(a.ID)
	roleService := Role_service.Role{}
	for _, v := range roleList {
		err := roleService.LoadPolicy(v)
		if err != nil {
			return err
		}
	}
	return nil
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
