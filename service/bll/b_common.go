package bll

import (
	"github.com/hequan2017/go-admin/service/menu_service"
	"github.com/hequan2017/go-admin/service/role_service"
	"github.com/hequan2017/go-admin/service/user_service"
)

type Common struct {
	UserAPI *user_service.User `inject:""`
	RoleAPI *Role_service.Role `inject:""`
	MenuAPI *menu_service.Menu `inject:""`
}
