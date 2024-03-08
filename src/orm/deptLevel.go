package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
)

func DeptLevel(appCtx *app.AppContext) *orm.ORMBase[*models.DeptLevel, *sqls.DeptLevel] {
	return &orm.ORMBase[*models.DeptLevel, *sqls.DeptLevel]{
		Executor: appCtx.DB,
		Stmt:     &sqls.DeptLevel{},
	}
}
