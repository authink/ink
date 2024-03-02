package orm

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/model"
	"github.com/authink/inkstone/orm"
	"github.com/jmoiron/sqlx"
)

type staff interface {
	orm.Inserter[models.Staff]
	orm.Saver[models.Staff]
	orm.Updater[models.Staff]
	orm.Geter[models.Staff]
	orm.Counter
	orm.Pager[models.Staff]
	GetByEmail(string) (*models.Staff, error)
}

type staffImpl app.AppContext

// Count implements staff.
func (s *staffImpl) Count(args ...any) (c int, err error) {
	err = s.DB.Get(&c, sqls.Staff.Count())
	return
}

// CountTx implements staff.
func (s *staffImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = tx.Get(&c, sqls.Staff.Count())
	return
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail(email string) (staff *models.Staff, err error) {
	staff = new(models.Staff)
	err = get(s.DB, staff, sqls.Staff.GetByEmail(), email)
	return
}

// Get implements staff.
// Subtle: this method shadows the method (*DB).Get of staffImpl.DB.
func (s *staffImpl) Get(id int) (staff *models.Staff, err error) {
	staff = new(models.Staff)
	err = get(s.DB, staff, sqls.Staff.Get(), id)
	return
}

// GetTx implements staff.
func (s *staffImpl) GetTx(tx *sqlx.Tx, id int) (staff *models.Staff, err error) {
	staff = new(models.Staff)
	err = get(tx, staff, sqls.Staff.GetForUpdate(), id)
	return
}

// Insert implements staff.
func (s *staffImpl) Insert(staff *models.Staff) error {
	return namedExec(s.DB, sqls.Staff.Insert(), staff, afterInsert)
}

// InsertTx implements staff.
func (s *staffImpl) InsertTx(tx *sqlx.Tx, staff *models.Staff) error {
	return namedExec(tx, sqls.Staff.Insert(), staff, afterInsert)
}

// PaginationTx implements staff.
func (s *staffImpl) PaginationTx(tx *sqlx.Tx, pager model.Pager) (staffs []models.Staff, err error) {
	stmt, err := tx.PrepareNamed(sqls.Staff.Pagination())
	if err != nil {
		return
	}
	err = stmt.Select(&staffs, pager)
	return
}

// Save implements staff.
func (s *staffImpl) Save(staff *models.Staff) error {
	return namedExec(s.DB, sqls.Staff.Save(), staff, afterSave)
}

// SaveTx implements staff.
func (s *staffImpl) SaveTx(tx *sqlx.Tx, staff *models.Staff) error {
	return namedExec(tx, sqls.Staff.Save(), staff, afterSave)
}

// Update implements staff.
func (s *staffImpl) Update(staff *models.Staff) error {
	return namedExec(s.DB, sqls.Staff.Update(), staff, nil)
}

// UpdateTx implements staff.
func (s *staffImpl) UpdateTx(tx *sqlx.Tx, staff *models.Staff) error {
	return namedExec(tx, sqls.Staff.Update(), staff, nil)
}

var _ staff = (*staffImpl)(nil)

func Staff(appCtx *app.AppContext) staff {
	return (*staffImpl)(appCtx)
}
