package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type AuthToken struct {
	sql.SQLBase
}

func (a *AuthToken) Count() string {
	return sbd.NewBuilder().
		Select(sql.Id.Count()).
		From(db.AuthToken.Tname()).
		String()
}

func (a *AuthToken) Delete() string {
	return sbd.NewBuilder().
		DeleteFrom(db.AuthToken.Tname()).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (a *AuthToken) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.AuthToken.Tname()).
		Columns(
			db.AuthToken.AccessToken,
			db.AuthToken.RefreshToken,
			db.AuthToken.AppId,
			db.AuthToken.AccountId,
		).
		String()
}

func (a *AuthToken) Pagination() string {
	aat := sbd.Table("at")
	aa := sbd.Table("a")
	return sbd.NewBuilder().
		Select(
			sql.Id.Of(aat),
			sql.CreatedAt.Of(aat),
			db.AuthToken.AccessToken.Of(aat),
			db.AuthToken.RefreshToken.Of(aat),
			db.AuthToken.AppId.Of(aat),
			db.App.Name.Of(aa).As(db.AuthTokenWithApp.AppName),
			db.AuthToken.AccountId.Of(aat),
		).
		From(
			db.AuthToken.Tname().As(aat),
			db.App.Tname().As(aa),
		).
		Where(sbd.Equal{
			Left:  db.AuthToken.AppId.Of(aat),
			Right: sql.Id.Of(aa),
		}).
		OrderBy(sql.Id.Of(aat)).
		Desc().
		Limit().
		String()
}

func (a *AuthToken) GetByAccessToken() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			db.AuthToken.AccessToken,
			db.AuthToken.RefreshToken,
			db.AuthToken.AppId,
			db.AuthToken.AccountId,
		).
		From(db.AuthToken.Tname()).
		Where(sbd.Equal{Left: db.AuthToken.AccessToken}).
		String()
}

func (a *AuthToken) GetByRefreshToken() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			db.AuthToken.AccessToken,
			db.AuthToken.RefreshToken,
			db.AuthToken.AppId,
			db.AuthToken.AccountId,
		).
		From(db.AuthToken.Tname()).
		Where(sbd.Equal{Left: db.AuthToken.RefreshToken}).
		String()
}
