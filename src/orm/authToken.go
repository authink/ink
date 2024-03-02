package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/model"
	"github.com/authink/inkstone/orm"
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
func (a *authTokenImpl) Count(...any) (c int, err error) {
	err = a.DB.Get(&c, sqls.AuthToken.Count())
	return
}

// CountTx implements authToken.
func (a *authTokenImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = tx.Get(&c, sqls.AuthToken.Count())
	return
}

// Delete implements authToken.
func (a *authTokenImpl) Delete(id int) (err error) {
	_, err = a.DB.Exec(
		sqls.AuthToken.Delete(),
		id,
	)
	return
}

// DeleteTx implements authToken.
func (a *authTokenImpl) DeleteTx(tx *sqlx.Tx, id int) (err error) {
	_, err = a.DB.Exec(
		sqls.AuthToken.Delete(),
		id,
	)
	return
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken(accessToken string) (token *models.AuthToken, err error) {
	token = new(models.AuthToken)
	err = get(a.DB, token, sqls.AuthToken.GetByAccessToken(), accessToken)
	return
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken(refreshToken string) (token *models.AuthToken, err error) {
	token = new(models.AuthToken)
	err = get(a.DB, token, sqls.AuthToken.GetByRefreshToken(), refreshToken)
	return
}

// Insert implements authToken.
func (a *authTokenImpl) Insert(token *models.AuthToken) error {
	return namedExec(a.DB, sqls.AuthToken.Insert(), token, afterInsert)
}

// InsertTx implements authToken.
func (a *authTokenImpl) InsertTx(tx *sqlx.Tx, token *models.AuthToken) error {
	return namedExec(tx, sqls.AuthToken.Insert(), token, afterInsert)
}

// PaginationTx implements authToken.
func (a *authTokenImpl) PaginationTx(tx *sqlx.Tx, pager model.Pager) (tokens []models.AuthTokenWithApp, err error) {
	stmt, err := tx.PrepareNamed(sqls.AuthToken.Pagination())
	if err != nil {
		return
	}
	err = stmt.Select(&tokens, pager)
	return
}

var _ authToken = (*authTokenImpl)(nil)

func AuthToken(appCtx *app.AppContext) authToken {
	return (*authTokenImpl)(appCtx)
}
