package sqls

import (
	"github.com/authink/ink/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type Staff struct {
	sql.SQLBase
}

func (s *Staff) Count() string {
	return sbd.NewBuilder().
		Select(sql.Id.Count()).
		From(db.Staff.Tname()).
		String()
}

func (s *Staff) Get() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			db.Staff.Email,
			db.Staff.Password,
			db.Staff.Active,
			db.Staff.Departure,
			db.Staff.Super,
			db.Staff.Phone,
		).
		From(db.Staff.Tname()).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (s *Staff) GetForUpdate() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			db.Staff.Email,
			db.Staff.Password,
			db.Staff.Active,
			db.Staff.Departure,
			db.Staff.Super,
			db.Staff.Phone,
		).
		From(db.Staff.Tname()).
		Where(sbd.Equal{Left: sql.Id}).
		ForUpdate().
		String()
}

func (s *Staff) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.Staff.Tname()).
		Columns(
			db.Staff.Email,
			db.Staff.Password,
			db.Staff.Phone,
			db.Staff.Super,
		).
		String()
}

func (s *Staff) Pagination() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			sql.CreatedAt,
			sql.UpdatedAt,
			db.Staff.Email,
			db.Staff.Phone,
			db.Staff.Super,
			db.Staff.Active,
			db.Staff.Departure,
		).
		From(db.Staff.Tname()).
		OrderBy(sql.Id).
		Desc().
		Limit().
		String()
}

func (s *Staff) Update() string {
	return sbd.NewBuilder().
		Update(db.Staff.Tname()).
		Set(
			db.Staff.Password,
			db.Staff.Phone,
			db.Staff.Super,
			db.Staff.Active,
			db.Staff.Departure,
		).
		Where(sbd.Equal{Left: sql.Id}).
		String()
}

func (s *Staff) GetByEmail() string {
	return sbd.NewBuilder().
		Select(
			sql.Id,
			db.Staff.Email,
			db.Staff.Password,
			db.Staff.Active,
			db.Staff.Departure,
			db.Staff.Super,
			db.Staff.Phone,
		).
		From(db.Staff.Tname()).
		Where(sbd.Equal{Left: db.Staff.Email}).
		String()
}

func (s *Staff) SelectStaffs() string {
	return sbd.NewBuilder().
		Select(sql.Id,
			db.Staff.Email,
		).
		From(db.Staff.Tname()).
		Where(&sbd.True{Left: db.Staff.Active}).
		And(&sbd.False{Left: db.Staff.Departure}).
		String()
}
