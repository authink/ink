package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
)

func App(appCtx *app.AppContext) *orm.ORMBase[*models.App, *sqls.App] {
	return &orm.ORMBase[*models.App, *sqls.App]{
		Executor: appCtx.DB,
		Stmt:     &sqls.App{},
	}
}
