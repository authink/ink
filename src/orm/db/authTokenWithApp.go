// Auto generated by inkstone, please do not change anything in this file
package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type authTokenWithApp struct {
	authToken

	AppName string
}

var AuthTokenWithApp authTokenWithApp

func init() {
	db.Bind(&AuthTokenWithApp, &models.AuthTokenWithApp{})
}