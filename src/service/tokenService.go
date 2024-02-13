package service

import (
	libsql "database/sql"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

type tokenService interface {
	SaveToken(*model.AuthToken) (libsql.Result, error)
	GetByRefreshToken(string) (*model.AuthToken, error)
	GetByAccessToken(string) (*model.AuthToken, error)
	DeleteToken(int) (libsql.Result, error)
}

type TokenService core.Ink

// GetByRefreshToken implements tokenService.
func (ts *TokenService) GetByRefreshToken(refreshToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = ts.DB.Get(
		token,
		sql.AuthToken.GetByRefreshToken(),
		refreshToken,
	)
	return
}

// GetByAccessToken implements tokenService.
func (ts *TokenService) GetByAccessToken(accessToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = ts.DB.Get(
		token,
		sql.AuthToken.GetByAccessToken(),
		accessToken,
	)
	return
}

// SaveToken implements tokenService.
func (ts *TokenService) SaveToken(token *model.AuthToken) (libsql.Result, error) {
	return ts.DB.NamedExec(
		sql.AuthToken.Insert(),
		token,
	)
}

// DeleteToken implements tokenService.
func (ts *TokenService) DeleteToken(id int) (libsql.Result, error) {
	return ts.DB.Exec(
		sql.AuthToken.Delete(),
		id,
	)
}

var _ tokenService = (*TokenService)(nil)
