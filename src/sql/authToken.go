package sql

import "fmt"

type authToken struct{}

// Delete implements SQL.
func (*authToken) Delete() string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = ?", table.AuthToken)
}

// Get implements SQL.
func (*authToken) Get() string {
	panic("unimplemented")
}

// Insert implements SQL.
func (*authToken) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (access_token, refresh_token, app_id, account_id) VALUES (:access_token, :refresh_token, :app_id, :account_id)", table.AuthToken)
}

func (*authToken) GetByRefreshToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE refresh_token = ?", table.AuthToken)
}

func (*authToken) GetByAccessToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE access_token = ?", table.AuthToken)
}

var _ SQL = (*authToken)(nil)
var AuthToken = &authToken{}
