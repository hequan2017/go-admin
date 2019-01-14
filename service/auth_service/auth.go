package auth_service

import (
	"github.com/hequan2017/go-admin/models"
)

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckUser(a.Username, a.Password)
}
