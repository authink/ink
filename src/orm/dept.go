package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/orm"
	"github.com/authink/stone/app"
)

func Dept(appCtx *app.AppContext) *orm.ORMBase[*models.Department, *sqls.Dept] {
	return &orm.ORMBase[*models.Department, *sqls.Dept]{
		Executor: appCtx.DB,
		Stmt:     &sqls.Dept{},
	}
}
