package orm

import (
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type staff interface {
	inkstone.ORM[model.Staff]
	GetByEmail(string) (*model.Staff, error)
}

type staffImpl inkstone.AppContext

// Delete implements staff.
func (*staffImpl) Delete(int) error {
	panic("unimplemented")
}

// Find implements staff.
func (*staffImpl) Find() ([]model.Staff, error) {
	panic("unimplemented")
}

// GetByEmail implements staff.
func (s *staffImpl) GetByEmail(email string) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = s.DB.Get(
		staff,
		sql.Staff.GetByEmail(),
		email,
	)
	return
}

// Get implements staff.
func (s *staffImpl) Get(id int) (staff *model.Staff, err error) {
	staff = &model.Staff{}
	err = s.DB.Get(
		staff,
		sql.Staff.Get(),
		id,
	)
	return
}

// Save implements staff.
func (s *staffImpl) Save(staff *model.Staff) (err error) {
	_, err = s.DB.NamedExec(
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

func Staff(app *inkstone.AppContext) staff {
	return (*staffImpl)(app)
}
