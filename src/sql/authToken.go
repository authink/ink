package sql

import (
	"fmt"

	"github.com/authink/inkstone"
)

type authToken struct{}

// Update implements inkstone.SQL.
func (*authToken) Update() string {
	panic("unimplemented")
}

// Find implements inkstone.SQL.
func (*authToken) Find() string {
	panic("unimplemented")
}

// Delete implements inkstone.SQL.
func (*authToken) Delete() string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = ?", table.AuthToken)
}

// Get implements inkstone.SQL.
func (*authToken) Get() string {
	panic("unimplemented")
}

// Insert implements inkstone.SQL.
func (*authToken) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (access_token, refresh_token, app_id, account_id) VALUES (:access_token, :refresh_token, :app_id, :account_id)", table.AuthToken)
}

func (*authToken) GetByRefreshToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE refresh_token = ?", table.AuthToken)
}

func (*authToken) GetByAccessToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE access_token = ?", table.AuthToken)
}

var _ inkstone.SQL = (*authToken)(nil)
var AuthToken = &authToken{}
