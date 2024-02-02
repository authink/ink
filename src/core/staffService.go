package core

import (
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

// GetStaffByEmail implements staffService.
func (ink *Ink) GetStaffByEmail(email string) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ink.db.Get(
		staff,
		sql.Staff.GetByEmail(),
		email,
	)
	return
}

// GetStaff implements staffService.
func (ink *Ink) GetStaff(id int) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ink.db.Get(
		staff,
		sql.Staff.Get(),
		id,
	)
	return
}

// SaveStaff implements staffService.
func (ink *Ink) SaveStaff(staff *model.Staff) (err error) {
	_, err = ink.db.NamedExec(
		sql.Staff.Insert(),
		staff,
	)
	return
}

// SaveStaffWithTx implements staffService.
func (*Ink) SaveStaffWithTx(staff *model.Staff, tx *sqlx.Tx) (err error) {
	_, err = tx.NamedExec(
		sql.Staff.Insert(),
		staff,
	)
	return
}

var _ staffService = (*Ink)(nil)
