package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
)

func DeptStaff(appCtx *app.AppContext) *orm.ORMBase[*models.DeptStaff, *sqls.DeptStaff] {
	return &orm.ORMBase[*models.DeptStaff, *sqls.DeptStaff]{
		Executor: appCtx.DB,
		Stmt:     &sqls.DeptStaff{},
	}
}
