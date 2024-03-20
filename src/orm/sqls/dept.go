package sqls

import (
	"github.com/authink/ink/src/orm/db"
	"github.com/authink/orm/sql"
	sbd "github.com/authink/sqlbuilder"
)

type Dept struct {
	sql.SQLBase
}

func (d *Dept) Insert() string {
	return sbd.NewBuilder().
		InsertInto(db.Department.Tname()).
		Columns(
			db.Department.Name,
			db.Department.OwnerId,
		).
		String()
}

func (d *Dept) Update() string {
	return sbd.NewBuilder().
		Update(db.Department.Tname()).
		Set(
			db.Department.Name,
			db.Department.OwnerId,
		).
		Where(&sbd.Equal{Left: sql.Id}).
		String()
}

func (d *Dept) Unique() string {
	return sbd.NewBuilder().
		Select(sql.Id.Count()).
		From(db.Department.Tname()).
		Where(&sbd.Equal{Left: db.Department.Name}).
		String()
}

func (d *Dept) Find() string {
	ad := sbd.Table("d")
	as := sbd.Table("s")
	return sbd.NewBuilder().
		Select(
			sql.Id.Of(ad),
			db.Department.Name.Of(ad),
			db.Department.OwnerId.Of(ad),
			db.Department.Active.Of(ad),
			db.Staff.Email.Of(as).As(db.DepartmentWithOwner.OwnerEmail),
		).
		From(
			db.Department.Tname().As(ad),
			db.Staff.Tname().As(as),
		).
		Where(sbd.Equal{
			Left:  db.Department.OwnerId.Of(ad),
			Right: sql.Id.Of(as),
		}).
		String()
}
