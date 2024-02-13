package service

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/jmoiron/sqlx"
)

type staffService interface {
	SaveStaff(*model.Staff) error
	SaveStaffWithTx(*model.Staff, *sqlx.Tx) error
	GetStaff(int) (*model.Staff, error)
	GetStaffByEmail(string) (*model.Staff, error)
}

type StaffService core.Ink

// GetStaffByEmail implements staffService.
func (ss *StaffService) GetStaffByEmail(email string) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ss.DB.Get(
		staff,
		sql.Staff.GetByEmail(),
		email,
	)
	return
}

// GetStaff implements staffService.
func (ss *StaffService) GetStaff(id int) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ss.DB.Get(
		staff,
		sql.Staff.Get(),
		id,
	)
	return
}

// SaveStaff implements staffService.
func (ss *StaffService) SaveStaff(staff *model.Staff) (err error) {
	_, err = ss.DB.NamedExec(
		sql.Staff.Insert(),
		staff,
	)
	return
}

// SaveStaffWithTx implements staffService.
func (*StaffService) SaveStaffWithTx(staff *model.Staff, tx *sqlx.Tx) (err error) {
	_, err = tx.NamedExec(
		sql.Staff.Insert(),
		staff,
	)
	return
}

var _ staffService = (*StaffService)(nil)
