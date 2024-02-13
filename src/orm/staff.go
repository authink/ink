package orm

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/jmoiron/sqlx"
)

type staff interface {
	Save(*model.Staff) error
	SaveWithTx(*model.Staff, *sqlx.Tx) error
	Get(int) (*model.Staff, error)
	GetByEmail(string) (*model.Staff, error)
}

type staffImpl core.Ink

// GetByEmail implements staff.
func (ss *staffImpl) GetByEmail(email string) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ss.DB.Get(
		staff,
		sql.Staff.GetByEmail(),
		email,
	)
	return
}

// Get implements staff.
func (ss *staffImpl) Get(id int) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = ss.DB.Get(
		staff,
		sql.Staff.Get(),
		id,
	)
	return
}

// Save implements staff.
func (ss *staffImpl) Save(staff *model.Staff) (err error) {
	_, err = ss.DB.NamedExec(
		sql.Staff.Insert(),
		staff,
	)
	return
}

// SaveWithTx implements staff.
func (*staffImpl) SaveWithTx(staff *model.Staff, tx *sqlx.Tx) (err error) {
	_, err = tx.NamedExec(
		sql.Staff.Insert(),
		staff,
	)
	return
}

var _ staff = (*staffImpl)(nil)

func Staff(ink *core.Ink) staff {
	return (*staffImpl)(ink)
}
