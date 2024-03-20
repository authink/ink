package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/orm"
	"github.com/authink/orm/model"
	"github.com/authink/stone/app"
)

type dept struct {
	*orm.ORMBase[*models.Department, *sqls.Dept]
}

func (d *dept) Unique(check *models.CheckUnique) (bool, error) {
	var c int
	err := orm.Count(d.Executor, d.Stmt.Unique(), &c, check)
	return c == 0, err
}

func (d *dept) Find() (depts []models.DepartmentWithOwner, err error) {
	err = orm.Select(d.Executor, d.Stmt.Find(), &depts, &model.Argument{})
	return
}

func Dept(appCtx *app.AppContext) *dept {
	return &dept{
		&orm.ORMBase[*models.Department, *sqls.Dept]{
			Executor: appCtx.DB,
			Stmt:     &sqls.Dept{},
		},
	}
}
