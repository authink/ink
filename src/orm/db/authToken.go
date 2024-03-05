package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type authToken struct {
	AccessToken  string
	RefreshToken string
	AppId        string
	AccountId    string
}

// Tname implements db.Table.
func (a *authToken) Tname() string {
	return "s_auth_tokens"
}

var _ db.Table = (*authToken)(nil)

var AuthToken authToken

type authTokenWithApp struct {
	authToken
	AppName string
}

var AuthTokenWithApp authTokenWithApp

func init() {
	db.Bind(&AuthToken, &models.AuthToken{})
	db.Bind(&AuthTokenWithApp, &models.AuthTokenWithApp{})
}
