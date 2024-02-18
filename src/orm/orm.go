package orm

import (
	"github.com/authink/ink.go/src/core"
	"github.com/jmoiron/sqlx"
)

type TxFunc func(tx *sqlx.Tx) error

type ORM[T any] interface {
	Save(*T) error
	SaveWithTx(*T, *sqlx.Tx) error
	Get(int) (*T, error)
	Find() ([]T, error)
	Delete(int) error
}

func Transaction(ink *core.Ink, txFunc TxFunc) (err error) {
	tx := ink.DB.MustBegin()

	if err = txFunc(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
