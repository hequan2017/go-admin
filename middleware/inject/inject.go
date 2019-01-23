package inject

import (
	"github.com/casbin/casbin"
	"github.com/facebookgo/inject"
	"github.com/hequan2017/go-admin/service/bll"
	"runtime"
)

// Object 注入对象
type Object struct {
	Common   *bll.Common
	Enforcer *casbin.Enforcer
}

// Init 初始化依赖注入
func Init() *Object {
	g := new(inject.Graph)

	// 注入casbin
	osType := runtime.GOOS
	var path string
	if osType == "windows"{
		path = "conf\\rbac_model.conf"
	}else if osType == "linux"{
		path = "conf/rbac_model.conf"
	}
	enforcer := casbin.NewEnforcer(path, false)
	_ = g.Provide(&inject.Object{Value: enforcer})

	Common := new(bll.Common)
	_ = g.Provide(&inject.Object{Value: Common})

	if err := g.Populate(); err != nil {
		panic("初始化依赖注入发生错误：" + err.Error())
	}

	return &Object{
		Enforcer: enforcer,
		Common:   Common,
	}
}
