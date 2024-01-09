package main

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sql"
)

func seed(ink *Ink) (err error) {
	tx := ink.db.MustBegin()
	defer txEnd(tx, err)

	if _, err = tx.NamedExec(
		sql.Query.InsertApp,
		models.NewApp(
			"admin.dev",
			"123456",
		),
	); err != nil {
		return
	}

	return
}
