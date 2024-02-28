package authz

import (
	_ "embed"

	"github.com/authink/inkstone"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/memwey/casbin-sqlx-adapter"
)

const (
	ResourceApp   string = "apps"
	ResourceToken string = "tokens"
	ResourceStaff string = "staffs"
	ResourceGroup string = "groups"
)

var (
	//go:embed rbac.conf
	rbac     string
	enforcer *casbin.Enforcer
)

func SetupEnforcer(appCtx *inkstone.AppContext) {
	model, err := model.NewModelFromString(rbac)
	if err != nil {
		panic(err)
	}

	adapter := sqlxadapter.NewAdapterByDB(appCtx.DB)

	enforcer, err = casbin.NewEnforcer(model, adapter)
	if err != nil {
		panic(err)
	}

	if err = enforcer.LoadPolicy(); err != nil {
		panic(err)
	}
}

func RBACEnforcer() *casbin.Enforcer {
	return enforcer
}
