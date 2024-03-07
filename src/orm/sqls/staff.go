package sqls

import (
	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
	"github.com/huandu/go-sqlbuilder"
)

type staff interface {
	sql.Inserter
	sql.Updater
	sql.Geter
	sql.GeterForUpdate
	sql.Counter
	sql.Pager
	GetByEmail() string
}

type staffImpl struct{}

// Count implements staff.
func (s *staffImpl) Count() (statement string) {
	tbnAlias := "s"
	sb := sqlbuilder.NewSelectBuilder()
	statement, _ = sb.
		Select(sql.Count(tbnAlias)).
		From(sb.As(db.Staff.Tname(), tbnAlias)).
		Build()
	return statement
}

// Get implements staff.
func (s *staffImpl) Get() (statement string) {
	statement, _ = sqlbuilder.
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
		Where(sql.EQ(sql.Id, "?")).
		Build()
	return statement
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail() (statement string) {
	statement, _ = sqlbuilder.
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
		Where(sql.EQ(db.Staff.Email, "?")).
		Build()
	return statement
}

// GetForUpdate implements staff.
func (s *staffImpl) GetForUpdate() (statement string) {
	statement, _ = sqlbuilder.
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
		Where(sql.EQ(sql.Id, "?")).
		ForUpdate().
		Build()
	return statement
}

// Insert implements staff.
func (s *staffImpl) Insert() (statement string) {
	statement, _ = sqlbuilder.
		InsertInto(db.Staff.Tname()).
		Cols(
			db.Staff.Email,
			db.Staff.Password,
			db.Staff.Phone,
			db.Staff.Super,
		).
		Values(
			sql.Named(db.Staff.Email),
			sql.Named(db.Staff.Password),
			sql.Named(db.Staff.Phone),
			sql.Named(db.Staff.Super),
		).Build()
	return sql.ReplaceAtWithColon(statement)
}

// Pagination implements staff.
func (s *staffImpl) Pagination() (statement string) {
	statement, _ = sqlbuilder.Select(
		sql.Id,
		sql.CreatedAt,
		sql.UpdatedAt,
		db.Staff.Email,
		db.Staff.Phone,
		db.Staff.Super,
		db.Staff.Active,
		db.Staff.Departure,
	).From(db.Staff.Tname()).
		OrderBy(sql.Id).
		Desc().
		Build()
	return sql.LimitAndOffset(statement)
}

// Update implements staff.
func (s *staffImpl) Update() (statement string) {
	sb := sqlbuilder.NewUpdateBuilder()
	statement, _ = sb.
		Update(db.Staff.Tname()).
		Set(
			sb.Assign(db.Staff.Password, sql.Named(db.Staff.Password)),
			sb.Assign(db.Staff.Phone, sql.Named(db.Staff.Phone)),
			sb.Assign(db.Staff.Super, sql.Named(db.Staff.Super)),
			sb.Assign(db.Staff.Active, sql.Named(db.Staff.Active)),
			sb.Assign(db.Staff.Departure, sql.Named(db.Staff.Departure)),
		).
		Where(sb.EQ(sql.Id, sql.Named(sql.Id))).
		Build()
	return sql.ReplaceAtWithColon(statement)
}

var Staff staff = &staffImpl{}
