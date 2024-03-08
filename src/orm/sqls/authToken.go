package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
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
	return sbd.NewBuilder().
		Select(sbd.Field(sql.Id).Count().As("c")).
		From(sbd.Table(db.AuthToken.Tname())).
		String()
}

// Delete implements authToken.
func (a *authTokenImpl) Delete() string {
	return sbd.NewBuilder().
		DeleteFrom(sbd.Table(db.AuthToken.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

// GetByAccessToken implements authToken.
func (a *authTokenImpl) GetByAccessToken() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			sbd.Field(db.AuthToken.AccessToken),
			sbd.Field(db.AuthToken.RefreshToken),
			sbd.Field(db.AuthToken.AppId),
			sbd.Field(db.AuthToken.AccountId),
		).
		From(sbd.Table(db.AuthToken.Tname())).
		Where(sbd.Equal{Left: sbd.Field(db.AuthToken.AccessToken)}).
		String()
}

// GetByRefreshToken implements authToken.
func (a *authTokenImpl) GetByRefreshToken() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			sbd.Field(db.AuthToken.AccessToken),
			sbd.Field(db.AuthToken.RefreshToken),
			sbd.Field(db.AuthToken.AppId),
			sbd.Field(db.AuthToken.AccountId),
		).
		From(sbd.Table(db.AuthToken.Tname())).
		Where(sbd.Equal{Left: sbd.Field(db.AuthToken.RefreshToken)}).
		String()
}

// Insert implements authToken.
func (a *authTokenImpl) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.AuthToken.Tname())).
		Columns(
			sbd.Field(db.AuthToken.AccessToken),
			sbd.Field(db.AuthToken.RefreshToken),
			sbd.Field(db.AuthToken.AppId),
			sbd.Field(db.AuthToken.AccountId),
		).
		String()
}

// Pagination implements authToken.
func (a *authTokenImpl) Pagination() string {
	aat := "at"
	aa := "a"
	return sbd.NewBuilder().
		Select(
			sbd.Field(sql.Id).Of(aat),
			sbd.Field(sql.CreatedAt).Of(aat),
			sbd.Field(db.AuthToken.AccessToken).Of(aat),
			sbd.Field(db.AuthToken.RefreshToken).Of(aat),
			sbd.Field(db.AuthToken.AppId).Of(aat),
			sbd.Field(db.App.Name).Of(aa).As(db.AuthTokenWithApp.AppName),
			sbd.Field(db.AuthToken.AccountId).Of(aat),
		).
		From(
			sbd.Table(db.AuthToken.Tname()).As(aat),
			sbd.Table(db.App.Tname()).As(aa),
		).
		Where(sbd.Equal{
			Left:  sbd.Field(db.AuthToken.AppId).Of(aat),
			Right: sbd.Field(sql.Id).Of(aa),
		}).
		OrderBy(sbd.Field(sql.Id).Of(aat)).
		Desc().
		Limit().
		String()
}

var AuthToken authToken = &authTokenImpl{}
