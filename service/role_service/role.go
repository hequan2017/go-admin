package Role_service

import (
	"github.com/casbin/casbin"
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

	Enforcer *casbin.Enforcer `inject:""`
}

func (a *Role) Add() error {
	role := map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	}
	if err := models.AddRole(role); err != nil {
		return err
	}

	return a.LoadPolicy(a.ID)
}

func (a *Role) Edit() error {
	err := models.EditRole(a.ID, map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	})
	if err != nil {
		return err
	}
	return a.LoadPolicy(a.ID)
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
	err := models.DeleteRole(a.ID)
	if err != nil {
		return err
	}
	a.Enforcer.DeletePermissionsForUser(a.Name)
	return nil
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

// LoadAllPolicy 加载所有的角色策略
func (a *Role) LoadAllPolicy() error {
	roles, err := models.GetRolesAll()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	for _, role := range roles {
		err = a.LoadPolicy(role.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadPolicy 加载角色权限策略
func (a *Role) LoadPolicy(id int) error {

	role, err := models.GetRole(id)
	if err != nil {
		return err
	}
	a.Enforcer.DeletePermissionsForUser(role.Name)
	for _, menu := range role.Menu {
		if menu.Path == "" || menu.Method == "" {
			continue
		}
		a.Enforcer.AddPermissionForUser(role.Name, menu.Path, menu.Method)
	}
	return nil
}
