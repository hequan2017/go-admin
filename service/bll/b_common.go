package bll

import (
	"go-admin/service/menu_service"
	"go-admin/service/role_service"
	"go-admin/service/user_service"
)

type Common struct {
	UserAPI *user_service.User `inject:""`
	RoleAPI *Role_service.Role `inject:""`
	MenuAPI *menu_service.Menu `inject:""`
}
