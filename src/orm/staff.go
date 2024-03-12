package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/orm"
	"github.com/authink/stone/app"
)

type staff struct {
	*orm.ORMBase[*models.Staff, *sqls.Staff]
}

func (s *staff) GetByEmail(staff *models.Staff) error {
	return orm.Get(s.Executor, s.Stmt.GetByEmail(), staff)
}

func Staff(appCtx *app.AppContext) *staff {
	return &staff{
		&orm.ORMBase[*models.Staff, *sqls.Staff]{
			Executor: appCtx.DB,
			Stmt:     &sqls.Staff{},
		},
	}
}
