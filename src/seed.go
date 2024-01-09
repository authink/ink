package main

import (
	"github.com/authink/ink.go/src/models"
	"github.com/authink/ink.go/src/sql"
)

func seed(ink *Ink) (err error) {
	admin, err := models.NewStaff(
		"admin@huoyijie.cn", "123456",
		"18222222222",
		true,
	)
	if err != nil {
		return
	}

	tx := ink.db.MustBegin()
	defer txEnd(tx, err)

	if _, err = tx.NamedExec(
		sql.Query.InsertStaff,
		admin,
	); err != nil {
		return
	}

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
