package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/tables"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
)

var (
	tbToken  = tables.AuthToken
	tbnToken = tbToken.TbName()
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
	return fmt.Sprintf("SELECT COUNT(id) c FROM %s", tbnToken)
}

// Delete implements authToken.
func (a *authTokenImpl) Delete() string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = ?", tbnToken)
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE access_token = ?", tbnToken)
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE refresh_token = ?", tbnToken)
}

// Insert implements authToken.
func (a *authTokenImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (access_token, refresh_token, app_id, account_id) VALUES (:access_token, :refresh_token, :app_id, :account_id)", tbnToken)
}

// Pagination implements authToken.
func (a *authTokenImpl) Pagination() (statement string) {
	tbnAlias1 := "at"
	tbnAlias2 := "a"
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sqlbuilder.
		Select(
			sql.Col(tbnAlias1, sql.Id),
			sql.Col(tbnAlias1, sql.CreatedAt),
			sql.Col(tbnAlias1, tbToken.AccessToken),
			sql.Col(tbnAlias1, tbToken.RefreshToken),
			sql.Col(tbnAlias1, tbToken.AppId),
			sb.As(
				sql.Col(tbnAlias2, tbApp.Name),
				"app_name",
			),
			sql.Col(tbnAlias1, tbToken.AccountId),
		).
		From(
			sb.As(tbnToken, tbnAlias1),
			sb.As(tbnApp, tbnAlias2),
		).
		Where(sql.EQ(
			sql.Col(tbnAlias1, tbToken.AppId),
			sql.Col(tbnAlias2, sql.Id),
		)).
		OrderBy(sql.Col(tbnAlias1, sql.Id)).
		Desc().
		Build()
	return sql.LimitAndOffset(statement)
}

var AuthToken authToken = new(authTokenImpl)
