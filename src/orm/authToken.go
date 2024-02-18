package orm

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

type authToken interface {
	Save(*model.AuthToken) error
	GetByRefreshToken(string) (*model.AuthToken, error)
	GetByAccessToken(string) (*model.AuthToken, error)
	Delete(int) error
}

type authTokenImpl core.Ink

// GetByRefreshToken implements authToken.
func (at *authTokenImpl) GetByRefreshToken(refreshToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = at.DB.Get(
		token,
		sql.AuthToken.GetByRefreshToken(),
		refreshToken,
	)
	return
}

// GetByAccessToken implements authToken.
func (at *authTokenImpl) GetByAccessToken(accessToken string) (token *model.AuthToken, err error) {
	token = &model.AuthToken{}
	err = at.DB.Get(
		token,
		sql.AuthToken.GetByAccessToken(),
		accessToken,
	)
	return
}

// Save implements authToken.
func (at *authTokenImpl) Save(token *model.AuthToken) (err error) {
	_, err = at.DB.NamedExec(
		sql.AuthToken.Insert(),
		token,
	)
	return
}

// Delete implements authToken.
func (at *authTokenImpl) Delete(id int) (err error) {
	_, err = at.DB.Exec(
		sql.AuthToken.Delete(),
		id,
	)
	return
}

var _ authToken = (*authTokenImpl)(nil)

func AuthToken(ink *core.Ink) authToken {
	return (*authTokenImpl)(ink)
}
