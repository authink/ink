package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	a "github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type staff interface {
	orm.Inserter[model.Staff]
	orm.Saver[model.Staff]
	orm.Updater[model.Staff]
	orm.Geter[model.Staff]
	orm.Counter
	orm.Pager[model.Staff]
	GetByEmail(string) (*model.Staff, error)
}

type staffImpl a.AppContext

// Count implements staff.
func (s *staffImpl) Count(args ...any) (c int, err error) {
	err = s.DB.Get(&c, sql.Staff.Count())
	return
}

// CountTx implements staff.
func (s *staffImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = tx.Get(&c, sql.Staff.Count())
	return
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail(email string) (staff *model.Staff, err error) {
	staff = new(model.Staff)
	err = s.DB.Get(
		staff,
		sql.Staff.GetByEmail(),
		email,
	)
	return
}

// Get implements staff.
// Subtle: this method shadows the method (*DB).Get of staffImpl.DB.
func (s *staffImpl) Get(id int) (staff *model.Staff, err error) {
	staff = new(model.Staff)
	err = s.DB.Get(
		staff,
		sql.Staff.Get(),
		id,
	)
	return
}

// GetTx implements staff.
func (s *staffImpl) GetTx(tx *sqlx.Tx, id int) (staff *model.Staff, err error) {
	staff = new(model.Staff)
	err = tx.Get(
		staff,
		sql.Staff.GetForUpdate(),
		id,
	)
	return
}

// Insert implements staff.
func (s *staffImpl) Insert(staff *model.Staff) error {
	return namedExec(s.DB, sql.Staff.Insert(), staff, handleInsertResult)
}

// InsertTx implements staff.
func (s *staffImpl) InsertTx(tx *sqlx.Tx, staff *model.Staff) error {
	return namedExec(tx, sql.Staff.Insert(), staff, handleInsertResult)
}

// PaginationTx implements staff.
func (s *staffImpl) PaginationTx(tx *sqlx.Tx, page orm.Page) (staffs []model.Staff, err error) {
	stmt, err := tx.PrepareNamed(sql.Staff.Pagination())
	if err != nil {
		return
	}
	err = stmt.Select(&staffs, page)
	return
}

// Save implements staff.
func (s *staffImpl) Save(staff *model.Staff) error {
	return namedExec(s.DB, sql.Staff.Save(), staff, handleSaveResult)
}

// SaveTx implements staff.
func (s *staffImpl) SaveTx(tx *sqlx.Tx, staff *model.Staff) error {
	return namedExec(tx, sql.Staff.Save(), staff, handleSaveResult)
}

// Update implements staff.
func (s *staffImpl) Update(staff *model.Staff) error {
	return namedExec(s.DB, sql.Staff.Update(), staff, nil)
}

// UpdateTx implements staff.
func (s *staffImpl) UpdateTx(tx *sqlx.Tx, staff *model.Staff) error {
	return namedExec(tx, sql.Staff.Update(), staff, nil)
}

var _ staff = (*staffImpl)(nil)

func Staff(appCtx *a.AppContext) staff {
	return (*staffImpl)(appCtx)
}
