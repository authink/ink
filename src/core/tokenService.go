package core

import (
	libsql "database/sql"

	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

type tokenService interface {
	SaveToken(*model.AuthToken) (libsql.Result, error)
	GetByRefreshToken(refreshToken string) (*model.AuthToken, error)
	DeleteToken(id int) (libsql.Result, error)
}

// GetByRefreshToken implements tokenService.
func (ink *Ink) GetByRefreshToken(refreshToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = ink.DB.Get(
		token,
		sql.AuthToken.GetByRefreshToken(),
		refreshToken,
	)
	return
}

// SaveToken implements tokenService.
func (ink *Ink) SaveToken(token *model.AuthToken) (libsql.Result, error) {
	return ink.DB.NamedExec(
		sql.AuthToken.Insert(),
		token,
	)
}

// DeleteToken implements tokenService.
func (ink *Ink) DeleteToken(id int) (libsql.Result, error) {
	return ink.DB.Exec(
		sql.AuthToken.Delete(),
		id,
	)
}

var _ tokenService = (*Ink)(nil)
