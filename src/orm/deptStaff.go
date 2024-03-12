package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/orm"
	"github.com/authink/stone/app"
)

func DeptStaff(appCtx *app.AppContext) *orm.ORMBase[*models.DeptStaff, *sqls.DeptStaff] {
	return &orm.ORMBase[*models.DeptStaff, *sqls.DeptStaff]{
		Executor: appCtx.DB,
		Stmt:     &sqls.DeptStaff{},
	}
}
