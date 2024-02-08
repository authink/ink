package core

import (
	libsql "database/sql"

	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

type tokenService interface {
	SaveToken(*model.AuthToken) (libsql.Result, error)
	GetByRefreshToken(string) (*model.AuthToken, error)
	GetByAccessToken(string) (*model.AuthToken, error)
	DeleteToken(int) (libsql.Result, error)
}

// GetByRefreshToken implements tokenService.
func (ink *Ink) GetByRefreshToken(refreshToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = ink.db.Get(
		token,
		sql.AuthToken.GetByRefreshToken(),
		refreshToken,
	)
	return
}

// GetByAccessToken implements tokenService.
func (ink *Ink) GetByAccessToken(accessToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = ink.db.Get(
		token,
		sql.AuthToken.GetByAccessToken(),
		accessToken,
	)
	return
}

// SaveToken implements tokenService.
func (ink *Ink) SaveToken(token *model.AuthToken) (libsql.Result, error) {
	return ink.db.NamedExec(
		sql.AuthToken.Insert(),
		token,
	)
}

// DeleteToken implements tokenService.
func (ink *Ink) DeleteToken(id int) (libsql.Result, error) {
	return ink.db.Exec(
		sql.AuthToken.Delete(),
		id,
	)
}

var _ tokenService = (*Ink)(nil)
