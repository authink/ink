package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/authink/inkstone/orm/model"
	"github.com/jmoiron/sqlx"
)

type authToken interface {
	orm.Inserter[models.AuthToken]
	orm.Deleter[models.AuthToken]
	orm.Counter
	orm.Pager[models.AuthTokenWithApp]
	GetByAccessToken(string) (*models.AuthToken, error)
	GetByRefreshToken(string) (*models.AuthToken, error)
}

type authTokenImpl app.AppContext

// Count implements authToken.
func (a *authTokenImpl) Count(args ...any) (c int, err error) {
	err = orm.Get(a.DB, &c, sqls.AuthToken.Count(), args...)
	return
}

// CountTx implements authToken.
func (a *authTokenImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = orm.Get(tx, &c, sqls.AuthToken.Count(), args...)
	return
}

// Delete implements authToken.
func (a *authTokenImpl) Delete(id int) error {
	return orm.Delete(a.DB, sqls.AuthToken.Delete(), id)
}

// DeleteTx implements authToken.
func (a *authTokenImpl) DeleteTx(tx *sqlx.Tx, id int) error {
	return orm.Delete(tx, sqls.AuthToken.Delete(), id)
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken(accessToken string) (token *models.AuthToken, err error) {
	token = new(models.AuthToken)
	err = orm.Get(a.DB, token, sqls.AuthToken.GetByAccessToken(), accessToken)
	return
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken(refreshToken string) (token *models.AuthToken, err error) {
	token = new(models.AuthToken)
	err = orm.Get(a.DB, token, sqls.AuthToken.GetByRefreshToken(), refreshToken)
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
	err = orm.Pagination(tx, sqls.AuthToken.Pagination(), &tokens, pager)
	return
}

var _ authToken = (*authTokenImpl)(nil)

func AuthToken(appCtx *app.AppContext) authToken {
	return (*authTokenImpl)(appCtx)
}
