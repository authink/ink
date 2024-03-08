package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type Staff struct {
	sql.SQLBase
}

func (s *Staff) Count() string {
	return sbd.NewBuilder().
		Select(sbd.Field(sql.Id).Count()).
		From(sbd.Table(db.Staff.Tname())).
		String()
}

func (s *Staff) Get() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Phone),
		).
		From(sbd.Table(db.Staff.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (s *Staff) GetForUpdate() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Phone),
		).
		From(sbd.Table(db.Staff.Tname())).
		Where(sbd.Equal{Left: sql.Id}).
		ForUpdate().
		String()
}

func (s *Staff) Insert() string {
	return sbd.NewBuilder().
		InsertInto(sbd.Table(db.Staff.Tname())).
		Columns(
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Phone),
			sbd.Field(db.Staff.Super),
		).
		String()
}

func (s *Staff) Pagination() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			sql.UpdatedAt,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Phone),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
		).
		From(sbd.Table(db.Staff.Tname())).
		OrderBy(sql.Id).
		Desc().
		Limit().
		String()
}

func (s *Staff) Update() string {
	return sbd.NewBuilder().
		Update(sbd.Table(db.Staff.Tname())).
		Set(
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Phone),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (s *Staff) GetByEmail() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sbd.Field(db.Staff.Email),
			sbd.Field(db.Staff.Password),
			sbd.Field(db.Staff.Active),
			sbd.Field(db.Staff.Departure),
			sbd.Field(db.Staff.Super),
			sbd.Field(db.Staff.Phone),
		).
		From(sbd.Table(db.Staff.Tname())).
		Where(sbd.Equal{Left: sbd.Field(db.Staff.Email)}).
		String()
}
