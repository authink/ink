package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
	"github.com/authink/orm/model"
	"github.com/jmoiron/sqlx"
)

type authToken interface {
	orm.Inserter[models.AuthToken]
	orm.Deleter[models.AuthToken]
	orm.Counter
	orm.Pager[models.AuthTokenWithApp]
	GetByAccessToken(*models.AuthToken) error
	GetByRefreshToken(*models.AuthToken) error
}

type authTokenImpl app.AppContext

// Count implements authToken.
func (a *authTokenImpl) Count(...model.Arg) (c int, err error) {
	err = orm.Count(a.DB, sqls.AuthToken.Count(), &c, &model.Argument{})
	return
}

// CountTx implements authToken.
func (a *authTokenImpl) CountTx(tx *sqlx.Tx, args ...model.Arg) (c int, err error) {
	err = orm.Count(tx, sqls.AuthToken.Count(), &c, &model.Argument{})
	return
}

// Delete implements authToken.
func (a *authTokenImpl) Delete(token *models.AuthToken) error {
	return orm.Delete(a.DB, sqls.AuthToken.Delete(), token)
}

// DeleteTx implements authToken.
func (a *authTokenImpl) DeleteTx(tx *sqlx.Tx, token *models.AuthToken) error {
	return orm.Delete(tx, sqls.AuthToken.Delete(), token)
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken(token *models.AuthToken) (err error) {
	err = orm.Get(a.DB, sqls.AuthToken.GetByAccessToken(), token)
	return
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken(token *models.AuthToken) (err error) {
	err = orm.Get(a.DB, sqls.AuthToken.GetByRefreshToken(), token)
	return
}

// Insert implements authToken.
func (a *authTokenImpl) Insert(token *models.AuthToken) error {
	return orm.NamedInsert(a.DB, sqls.AuthToken.Insert(), token)
}

// InsertTx implements authToken.
func (a *authTokenImpl) InsertTx(tx *sqlx.Tx, token *models.AuthToken) error {
	return orm.NamedInsert(tx, sqls.AuthToken.Insert(), token)
}

// PaginationTx implements authToken.
func (a *authTokenImpl) PaginationTx(tx *sqlx.Tx, pager model.Pager) (tokens []models.AuthTokenWithApp, err error) {
	err = orm.Select(tx, sqls.AuthToken.Pagination(), &tokens, pager)
	return
}

var _ authToken = (*authTokenImpl)(nil)

func AuthToken(appCtx *app.AppContext) authToken {
	return (*authTokenImpl)(appCtx)
}
