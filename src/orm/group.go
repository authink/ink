package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
	"github.com/authink/orm/model"
	"github.com/jmoiron/sqlx"
)

type group struct {
	*orm.ORMBase[*models.Group, *sqls.Group]
}

func (g *group) PaginationTx(tx *sqlx.Tx, pager model.Pager) (groups []models.GroupWithApp, err error) {
	err = orm.Select(tx, g.Stmt.Pagination(), &groups, pager)
	return
}

func Group(appCtx *app.AppContext) *group {
	return &group{
		&orm.ORMBase[*models.Group, *sqls.Group]{
			Executor: appCtx.DB,
			Stmt:     &sqls.Group{},
		},
	}
}
