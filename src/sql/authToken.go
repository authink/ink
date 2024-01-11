package sql

import "fmt"

type authToken struct{}

func (*authToken) GetByRefreshToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE refresh_token = ?", table.AuthToken)
}

// Get implements query.
func (*authToken) Get() string {
	panic("unimplemented")
}

// Insert implements query.
func (*authToken) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (access_token, refresh_token, app_id, account_id) VALUES (:access_token, :refresh_token, :app_id, :account_id)", table.AuthToken)
}

var _ query = (*authToken)(nil)
var AuthToken = &authToken{}
