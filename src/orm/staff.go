package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/inkstone/orm"
	"github.com/authink/inkstone/orm/model"
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
	err = orm.Get(s.DB, &c, sqls.Staff.Count(), args...)
	return
}

// CountTx implements staff.
func (s *staffImpl) CountTx(tx *sqlx.Tx, args ...any) (c int, err error) {
	err = orm.Get(tx, &c, sqls.Staff.Count(), args...)
	return
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail(email string) (staff *models.Staff, err error) {
	staff = &models.Staff{}
	err = orm.Get(s.DB, staff, sqls.Staff.GetByEmail(), email)
	return
}

// Get implements staff.
// Subtle: this method shadows the method (*DB).Get of staffImpl.DB.
func (s *staffImpl) Get(id int) (staff *models.Staff, err error) {
	staff = &models.Staff{}
	err = orm.Get(s.DB, staff, sqls.Staff.Get(), id)
	return
}

// GetTx implements staff.
func (s *staffImpl) GetTx(tx *sqlx.Tx, id int) (staff *models.Staff, err error) {
	staff = &models.Staff{}
	err = orm.Get(tx, staff, sqls.Staff.GetForUpdate(), id)
	return
}

// Insert implements staff.
func (s *staffImpl) Insert(staff *models.Staff) error {
	return orm.NamedInsert(s.DB, sqls.Staff.Insert(), staff)
}

// InsertTx implements staff.
func (s *staffImpl) InsertTx(tx *sqlx.Tx, staff *models.Staff) error {
	return orm.NamedInsert(tx, sqls.Staff.Insert(), staff)
}

// PaginationTx implements staff.
func (s *staffImpl) PaginationTx(tx *sqlx.Tx, pager model.Pager) (staffs []models.Staff, err error) {
	err = orm.Pagination(tx, sqls.Staff.Pagination(), &staffs, pager)
	return
}

// Save implements staff.
func (s *staffImpl) Save(staff *models.Staff) error {
	return orm.NamedSave(s.DB, sqls.Staff.Save(), staff)
}

// SaveTx implements staff.
func (s *staffImpl) SaveTx(tx *sqlx.Tx, staff *models.Staff) error {
	return orm.NamedSave(tx, sqls.Staff.Save(), staff)
}

// Update implements staff.
func (s *staffImpl) Update(staff *models.Staff) error {
	return orm.NamedUpdate(s.DB, sqls.Staff.Update(), staff)
}

// UpdateTx implements staff.
func (s *staffImpl) UpdateTx(tx *sqlx.Tx, staff *models.Staff) error {
	return orm.NamedUpdate(tx, sqls.Staff.Update(), staff)
}

var _ staff = (*staffImpl)(nil)

func Staff(appCtx *app.AppContext) staff {
	return (*staffImpl)(appCtx)
}
