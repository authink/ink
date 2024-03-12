package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/orm"
	"github.com/authink/stone/app"
)

func DeptLevel(appCtx *app.AppContext) *orm.ORMBase[*models.DeptLevel, *sqls.DeptLevel] {
	return &orm.ORMBase[*models.DeptLevel, *sqls.DeptLevel]{
		Executor: appCtx.DB,
		Stmt:     &sqls.DeptLevel{},
	}
}
