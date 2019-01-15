package user_service

import (
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

	return nil
}

func (a *User) Edit() error {
	return models.EditUser(a.ID, map[string]interface{}{
		"password":    a.Password,
		"role_id": a.Role,
	},
	)
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
	return models.DeleteUser(a.ID)
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
