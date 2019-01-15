package Role_service

import (
	"github.com/hequan2017/go-admin/models"
)

type Role struct {
	ID   int
	Name string
	Menu int

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int
}

func (a *Role) Add() error {
	role := map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	}
	if err := models.AddRole(role); err != nil {
		return err
	}

	return nil
}

func (a *Role) Edit() error {
	return models.EditRole(a.ID, map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	},
	)
}

func (a *Role) Get() (*models.Role, error) {

	Role, err := models.GetRole(a.ID)
	if err != nil {
		return nil, err
	}

	return Role, nil
}

func (a *Role) GetAll() ([]*models.Role, error) {
	Role, err := models.GetRoles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	return Role, nil
}

func (a *Role) Delete() error {
	return models.DeleteRole(a.ID)
}

func (a *Role) ExistByID() (bool, error) {
	return models.ExistRoleByID(a.ID)
}

func (a *Role) Count() (int, error) {
	return models.GetRoleTotal(a.getMaps())
}

func (a *Role) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}
