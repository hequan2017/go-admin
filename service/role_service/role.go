package Role_service

import (
	"errors"
	"github.com/casbin/casbin"
	"go-admin/models"
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

func (a *Role) Add() (id int, err error) {
	role := map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	}
	name, _ := models.CheckRoleName(a.Name)

	if name {
		return 0, errors.New("name 名字重复,请更改！")
	}

	if id, err := models.AddRole(role); err == nil {
		return id, nil
	} else {
		return 0, err
	}

}

func (a *Role) Edit() error {
	data := map[string]interface{}{
		"name":    a.Name,
		"menu_id": a.Menu,
	}
	name, _ := models.CheckRoleNameId(a.Name, a.ID)

	if name {
		return errors.New("name 名字重复,请更改！")
	}
	err := models.EditRole(a.ID, data)
	if err != nil {
		return err
	}
	if a.Menu != 0 {
		if err := a.LoadPolicy(a.ID); err != nil {
			return err
		}
	}
	return nil
}

func (a *Role) Get() (*models.Role, error) {

	Role, err := models.GetRole(a.ID)
	if err != nil {
		return nil, err
	}

	return Role, nil
}

func (a *Role) GetAll() ([]*models.Role, error) {
	if a.ID != 0 {
		maps := make(map[string]interface{})
		maps["deleted_on"] = 0
		maps["id"] = a.ID
		Role, err := models.GetRoles(a.PageNum, a.PageSize, maps)
		if err != nil {
			return nil, err
		}

		return Role, nil
	} else {
		Role, err := models.GetRoles(a.PageNum, a.PageSize, a.getMaps())
		if err != nil {
			return nil, err
		}
		return Role, nil
	}
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
	a.Enforcer.DeleteRole(role.Name)

	for _, menu := range role.Menu {
		if menu.Path == "" || menu.Method == "" {
			continue
		}
		a.Enforcer.AddPermissionForUser(role.Name, menu.Path, menu.Method)
	}
	return nil
}
