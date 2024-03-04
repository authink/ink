package tables

import "github.com/authink/inkstone/orm/sql"

type authToken struct {
	AccessToken  string
	RefreshToken string
	AppId        string
	AccountId    string
}

// TbName implements sql.Table.
func (a *authToken) TbName() string {
	return TB_TOKENS
}

var _ sql.Table = (*authToken)(nil)
