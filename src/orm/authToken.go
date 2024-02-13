package orm

import (
	libsql "database/sql"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

type authToken interface {
	Save(*model.AuthToken) (libsql.Result, error)
	GetByRefreshToken(string) (*model.AuthToken, error)
	GetByAccessToken(string) (*model.AuthToken, error)
	Delete(int) (libsql.Result, error)
}

type authTokenImpl core.Ink

// GetByRefreshToken implements authToken.
func (ts *authTokenImpl) GetByRefreshToken(refreshToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = ts.DB.Get(
		token,
		sql.AuthToken.GetByRefreshToken(),
		refreshToken,
	)
	return
}

// GetByAccessToken implements authToken.
func (ts *authTokenImpl) GetByAccessToken(accessToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = ts.DB.Get(
		token,
		sql.AuthToken.GetByAccessToken(),
		accessToken,
	)
	return
}

// Save implements authToken.
func (ts *authTokenImpl) Save(token *model.AuthToken) (libsql.Result, error) {
	return ts.DB.NamedExec(
		sql.AuthToken.Insert(),
		token,
	)
}

// Delete implements authToken.
func (ts *authTokenImpl) Delete(id int) (libsql.Result, error) {
	return ts.DB.Exec(
		sql.AuthToken.Delete(),
		id,
	)
}

var _ authToken = (*authTokenImpl)(nil)

func AuthToken(ink *core.Ink) authToken {
	return (*authTokenImpl)(ink)
}
