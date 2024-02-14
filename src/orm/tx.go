package orm

import (
	"github.com/authink/ink.go/src/core"
	"github.com/jmoiron/sqlx"
)

type TxFunc func(tx *sqlx.Tx) error

func Transaction(ink *core.Ink, txFunc TxFunc) (err error) {
	tx := ink.DB.MustBegin()

	if err = txFunc(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
