package migrate

import (
	"log"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

func Seed(ink *core.Ink) {
	admin, err := model.NewStaff(
		"admin@huoyijie.cn", "123456",
		"18222222222",
		true,
	)
	if err != nil {
		log.Fatalf("Seed: %s\n", err)
	}

	tx := ink.DB.MustBegin()
	defer core.TxEnd(tx, err)

	if _, err = tx.NamedExec(
		sql.Query.InsertStaff,
		admin,
	); err != nil {
		log.Fatalf("Seed: %s\n", err)
	}

	if _, err = tx.NamedExec(
		sql.Query.InsertApp,
		model.NewApp(
			"admin.dev",
			"123456",
		),
	); err != nil {
		log.Fatalf("Seed: %s\n", err)
	}
}
