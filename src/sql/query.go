package sql

import "fmt"

type query struct {
	InsertApp                  string
	GetApp                     string
	InsertStaff                string
	GetStaff                   string
	InsertAuthToken            string
	GetAuthTokenByRefreshToken string
}

var Query = &query{
	InsertApp: fmt.Sprintf("INSERT INTO %s (name, secret) VALUES (:name, :secret)", Table.App),

	GetApp: fmt.Sprintf("SELECT id, name, secret, active FROM %s WHERE id = ?", Table.App),

	InsertStaff: fmt.Sprintf("INSERT INTO %s (email, password, phone, super) VALUES (:email, :password, :phone, :super)", Table.Staff),

	GetStaff: fmt.Sprintf("SELECT id, email, password, active, departure, super, phone FROM %s WHERE email = ?", Table.Staff),

	InsertAuthToken: fmt.Sprintf("INSERT INTO %s (access_token, refresh_token, app_id, account_id) VALUES (:access_token, :refresh_token, :app_id, :account_id)", Table.AuthToken),

	GetAuthTokenByRefreshToken: fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE refresh_token = ?", Table.AuthToken),
}
