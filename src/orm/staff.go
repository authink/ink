package orm

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/orm/sqls"
	"github.com/authink/inkstone/app"
	"github.com/authink/orm"
	"github.com/authink/orm/model"
	"github.com/jmoiron/sqlx"
)

type staff interface {
	orm.Inserter[models.Staff]
	orm.Updater[models.Staff]
	orm.Geter[models.Staff]
	orm.Counter
	orm.Pager[models.Staff]
	GetByEmail(*models.Staff) error
}

type staffImpl app.AppContext

// Count implements staff.
func (s *staffImpl) Count(...model.Arg) (c int, err error) {
	err = orm.Count(s.DB, sqls.Staff.Count(), &c, &model.Argument{})
	return
}

// CountTx implements staff.
func (s *staffImpl) CountTx(tx *sqlx.Tx, args ...model.Arg) (c int, err error) {
	err = orm.Count(tx, sqls.Staff.Count(), &c, &model.Argument{})
	return
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail(staff *models.Staff) (err error) {
	err = orm.Get(s.DB, sqls.Staff.GetByEmail(), staff)
	return
}

// Get implements staff.
// Subtle: this method shadows the method (*DB).Get of staffImpl.DB.
func (s *staffImpl) Get(staff *models.Staff) (err error) {
	err = orm.Get(s.DB, sqls.Staff.Get(), staff)
	return
}

// GetTx implements staff.
func (s *staffImpl) GetTx(tx *sqlx.Tx, staff *models.Staff) (err error) {
	err = orm.Get(tx, sqls.Staff.GetForUpdate(), staff)
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
	err = orm.Select(tx, sqls.Staff.Pagination(), &staffs, pager)
	return
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
