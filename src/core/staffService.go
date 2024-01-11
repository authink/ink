package core

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

type staffService interface {
	SaveStaff(*model.Staff) error
	GetStaff(id int) (*model.Staff, error)
	GetStaffByEmail(email string) (*model.Staff, error)
}

// GetStaffByEmail implements staffService.
func (ink *Ink) GetStaffByEmail(email string) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ink.DB.Get(
		staff,
		sql.Staff.GetByEmail(),
		email,
	)
	return
}

// GetStaff implements staffService.
func (ink *Ink) GetStaff(id int) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ink.DB.Get(
		staff,
		sql.Staff.Get(),
		id,
	)
	return
}

// SaveStaff implements staffService.
func (ink *Ink) SaveStaff(*model.Staff) error {
	panic("unimplemented")
}

var _ staffService = (*Ink)(nil)
