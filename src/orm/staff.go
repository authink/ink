package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
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
