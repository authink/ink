package orm

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
	"github.com/authink/orm/model"
	"github.com/jmoiron/sqlx"
)

type authToken struct {
	*orm.ORMBase[*models.AuthToken, *sqls.AuthToken]
}

func (a *authToken) GetByAccessToken(token *models.AuthToken) error {
	return orm.Get(a.Executor, a.Stmt.GetByAccessToken(), token)
}

func (a *authToken) GetByRefreshToken(token *models.AuthToken) error {
	return orm.Get(a.Executor, a.Stmt.GetByRefreshToken(), token)
}

func (a *authToken) PaginationTx(tx *sqlx.Tx, pager model.Pager) (tokens []models.AuthTokenWithApp, err error) {
	err = orm.Select(tx, a.Stmt.Pagination(), &tokens, pager)
	return
}

func AuthToken(appCtx *app.AppContext) *authToken {
	return &authToken{
		&orm.ORMBase[*models.AuthToken, *sqls.AuthToken]{
			Executor: appCtx.DB,
			Stmt:     &sqls.AuthToken{},
		},
	}
}
