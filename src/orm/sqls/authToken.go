package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
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
func (a *authTokenImpl) Count() (statement string) {
	tbnAlias := "at"
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
		Select(
			sql.Count(tbnAlias),
		).
		From(sb.As(db.AuthToken.Tname(), tbnAlias)).
		Build()
	return statement
}

// Delete implements authToken.
func (a *authTokenImpl) Delete() (statement string) {
	statement, _ = sqlbuilder.
		DeleteFrom(db.AuthToken.Tname()).
		Where(sql.EQ(sql.Id, "?")).
		Build()
	return statement
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE access_token = ?", db.AuthToken.Tname())
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken() string {
	return fmt.Sprintf("SELECT id, created_at, access_token, refresh_token, app_id, account_id FROM %s WHERE refresh_token = ?", db.AuthToken.Tname())
}

// Insert implements authToken.
func (a *authTokenImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (access_token, refresh_token, app_id, account_id) VALUES (:access_token, :refresh_token, :app_id, :account_id)", db.AuthToken.Tname())
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
			sql.Col(tbnAlias1, db.AuthToken.AccessToken),
			sql.Col(tbnAlias1, db.AuthToken.RefreshToken),
			sql.Col(tbnAlias1, db.AuthToken.AppId),
			sb.As(
				sql.Col(tbnAlias2, db.App.Name),
				db.AuthTokenWithApp.AppName,
			),
			sql.Col(tbnAlias1, db.AuthToken.AccountId),
		).
		From(
			sb.As(db.AuthToken.Tname(), tbnAlias1),
			sb.As(db.App.Tname(), tbnAlias2),
		).
		Where(sql.EQ(
			sql.Col(tbnAlias1, db.AuthToken.AppId),
			sql.Col(tbnAlias2, sql.Id),
		)).
		OrderBy(sql.Col(tbnAlias1, sql.Id)).
		Desc().
		Build()
	return sql.LimitAndOffset(statement)
}

var AuthToken authToken = &authTokenImpl{}
