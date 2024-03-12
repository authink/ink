package authz

import (
	_ "embed"

	"github.com/authink/stone/app"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	sqlxadapter "github.com/memwey/casbin-sqlx-adapter"
)

var (
	//go:embed rbac.conf
	rbac     string
	enforcer *casbin.Enforcer
)

func SetupEnforcer(appCtx *app.AppContext) {
	model, err := model.NewModelFromString(rbac)
	if err != nil {
		panic(err)
	}

	adapter := sqlxadapter.NewAdapterByDB(appCtx.DB.DB)

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
