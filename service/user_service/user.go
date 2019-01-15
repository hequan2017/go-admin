package user_service

import (
	"github.com/casbin/casbin"
	"github.com/hequan2017/go-admin/models"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     int

	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int

	Enforcer *casbin.Enforcer `inject:""`
}

func (a *User) Check() (bool, error) {
	return models.CheckUser(a.Username, a.Password)
}

func (a *User) Add() error {
	menu := map[string]interface{}{
		"username": a.Username,
		"password": a.Password,
		"role_id":  a.Role,
	}
	if err := models.AddUser(menu); err != nil {
		return err
	}

	return a.LoadPolicy(a.ID)
}

func (a *User) Edit() error {
	err := models.EditUser(a.ID, map[string]interface{}{
		"password": a.Password,
		"role_id":  a.Role,
	})
	if err != nil {
		return err
	}
	return a.LoadPolicy(a.ID)
}

func (a *User) Get() (*models.User, error) {

	user, err := models.GetUser(a.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *User) GetAll() ([]*models.User, error) {
	user, err := models.GetUsers(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a *User) Delete() error {
	err := models.DeleteUser(a.ID)
	if err != nil {
		return err
	}
	a.Enforcer.DeleteUser(a.Username)
	return nil
}

func (a *User) ExistByID() (bool, error) {
	return models.ExistUserByID(a.ID)
}

func (a *User) Count() (int, error) {
	return models.GetUserTotal(a.getMaps())
}

func (a *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	return maps
}

// LoadAllPolicy 加载所有的用户策略
func (a *User) LoadAllPolicy() error {
	users, err := models.GetUsersAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		err = a.LoadPolicy(user.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadPolicy 加载用户权限策略
func (a *User) LoadPolicy(id int) error {

	user, err := models.GetUser(id)
	if err != nil {
		return err
	}
	a.Enforcer.DeleteRolesForUser(user.Username)
	for _, ro := range user.Role {
		a.Enforcer.AddRoleForUser(user.Username, ro.Name)
	}
	return nil
}
