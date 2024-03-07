package sqls

import (
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
	sb := sqlbuilder.NewDeleteBuilder()
	statement, _ = sb.
		DeleteFrom(db.AuthToken.Tname()).
		Where(sb.EQ(sql.Id, sql.Named(sql.Id))).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken() (statement string) {
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
		Select(
			sql.Id,
			sql.CreatedAt,
			db.AuthToken.AccessToken,
			db.AuthToken.RefreshToken,
			db.AuthToken.AppId,
			db.AuthToken.AccountId,
		).
		From(db.AuthToken.Tname()).
		Where(sb.EQ(db.AuthToken.AccessToken, sql.Named(db.AuthToken.AccessToken))).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken() (statement string) {
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
		Select(
			sql.Id,
			sql.CreatedAt,
			db.AuthToken.AccessToken,
			db.AuthToken.RefreshToken,
			db.AuthToken.AppId,
			db.AuthToken.AccountId,
		).
		From(db.AuthToken.Tname()).
		Where(sb.EQ(db.AuthToken.RefreshToken, sql.Named(db.AuthToken.RefreshToken))).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

// Insert implements authToken.
func (a *authTokenImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.
		InsertInto(db.AuthToken.Tname()).
		Cols(
			db.AuthToken.AccessToken,
			db.AuthToken.RefreshToken,
			db.AuthToken.AppId,
			db.AuthToken.AccountId,
		).Values(
		sql.Named(db.AuthToken.AccessToken),
		sql.Named(db.AuthToken.RefreshToken),
		sql.Named(db.AuthToken.AppId),
		sql.Named(db.AuthToken.AccountId),
	).Build()
	return sql.ReplaceAtWithColon(statement)
}

// Pagination implements authToken.
func (a *authTokenImpl) Pagination() (statement string) {
	tbnAlias1 := "at"
	tbnAlias2 := "a"
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
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
