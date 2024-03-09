package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
)

func Log(appCtx *app.AppContext) *orm.ORMBase[*models.Log, *sqls.Log] {
	return &orm.ORMBase[*models.Log, *sqls.Log]{
		Executor: appCtx.DB,
		Stmt:     &sqls.Log{},
	}
}
