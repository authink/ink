package sqls

import (
	"fmt"

	"github.com/authink/inkstone/sql"
)

type authToken interface {
	sql.Inserter
	sql.Deleter
	sql.Counter
	sql.Pager
	GetByAccessToken() string
	GetByRefreshToken() string
}

type authTokenImpl struct{}

// Count implements authToken.
func (a *authTokenImpl) Count() string {
	return fmt.Sprintf("SELECT COUNT(id) c FROM %s", table.AuthToken)
}

// Delete implements authToken.
func (a *authTokenImpl) Delete() string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = ?", table.AuthToken)
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE access_token = ?", table.AuthToken)
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE refresh_token = ?", table.AuthToken)
}

// Insert implements authToken.
func (a *authTokenImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (access_token, refresh_token, app_id, account_id) VALUES (:access_token, :refresh_token, :app_id, :account_id)", table.AuthToken)
}

// Pagination implements authToken.
func (a *authTokenImpl) Pagination() string {
	return fmt.Sprintf("SELECT at.id, at.created_at, at.access_token, at.refresh_token, at.app_id, a.name app_name, at.account_id FROM %s at, %s a WHERE at.app_id = a.id ORDER BY at.id DESC LIMIT :limit OFFSET :offset", table.AuthToken, table.App)
}

var AuthToken authToken = new(authTokenImpl)
