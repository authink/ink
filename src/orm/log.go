package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/orm"
	"github.com/authink/stone/app"
)

func Log(appCtx *app.AppContext) *orm.ORMBase[*models.Log, *sqls.Log] {
	return &orm.ORMBase[*models.Log, *sqls.Log]{
		Executor: appCtx.DB,
		Stmt:     &sqls.Log{},
	}
}
