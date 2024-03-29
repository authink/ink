// Package db Code generated by authink/orm. DO NOT EDIT
package db

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/orm/db"
	sbd "github.com/authink/sqlbuilder"
)

type authToken struct {
	AccessToken sbd.Field

	RefreshToken sbd.Field

	AppId sbd.Field

	AccountId sbd.Field
}

// Tname implements db.Table.
func (*authToken) Tname() sbd.Table {
	return "s_auth_tokens"
}

var _ db.Table = (*authToken)(nil)

var AuthToken authToken

func init() {
	db.Bind(&AuthToken, &models.AuthToken{})
}
