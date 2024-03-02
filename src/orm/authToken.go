package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type authToken interface {
	orm.Inserter[model.AuthToken]
	orm.Deleter[model.AuthToken]
	orm.Counter
	orm.Pager[model.AuthTokenWithApp]
	GetByAccessToken(string) (*model.AuthToken, error)
	GetByRefreshToken(string) (*model.AuthToken, error)
}

type authTokenImpl a.AppContext

// Count implements authToken.
func (a *authTokenImpl) Count(...any) (c int, err error) {
	err = a.DB.Get(&c, sql.AuthToken.Count())
	return
}

// CountTx implements authToken.
func (a *authTokenImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = tx.Get(&c, sql.AuthToken.Count())
	return
}

// Delete implements authToken.
func (a *authTokenImpl) Delete(id int) (err error) {
	_, err = a.DB.Exec(
		sql.AuthToken.Delete(),
		id,
	)
	return
}

// DeleteTx implements authToken.
func (a *authTokenImpl) DeleteTx(tx *sqlx.Tx, id int) (err error) {
	_, err = a.DB.Exec(
		sql.AuthToken.Delete(),
		id,
	)
	return
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken(accessToken string) (token *model.AuthToken, err error) {
	token = new(model.AuthToken)
	err = a.DB.Get(
		token,
		sql.AuthToken.GetByAccessToken(),
		accessToken,
	)
	return
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken(refreshToken string) (token *model.AuthToken, err error) {
	token = new(model.AuthToken)
	err = a.DB.Get(
		token,
		sql.AuthToken.GetByRefreshToken(),
		refreshToken,
	)
	return
}

// Insert implements authToken.
func (a *authTokenImpl) Insert(token *model.AuthToken) error {
	return namedExec(a.DB, sql.AuthToken.Insert(), token, handleInsertResult)
}

// InsertTx implements authToken.
func (a *authTokenImpl) InsertTx(tx *sqlx.Tx, token *model.AuthToken) error {
	return namedExec(tx, sql.AuthToken.Insert(), token, handleInsertResult)
}

// PaginationTx implements authToken.
func (a *authTokenImpl) PaginationTx(tx *sqlx.Tx, page orm.Page) (tokens []model.AuthTokenWithApp, err error) {
	stmt, err := tx.PrepareNamed(sql.AuthToken.Pagination())
	if err != nil {
		return
	}
	err = stmt.Select(&tokens, page)
	return
}

var _ authToken = (*authTokenImpl)(nil)

func AuthToken(appCtx *a.AppContext) authToken {
	return (*authTokenImpl)(appCtx)
}
