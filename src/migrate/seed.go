package migrate

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
)

func Seed(ink *core.Ink) (err error) {
	admin, err := model.NewStaff(
		"admin@huoyijie.cn", "123456",
		"18222222222",
		true,
	)
	if err != nil {
		return
	}

	tx := ink.DB.MustBegin()
	defer core.TxEnd(tx, err)

	if _, err = tx.NamedExec(
		sql.Query.InsertStaff,
		admin,
	); err != nil {
		return
	}

	if _, err = tx.NamedExec(
		sql.Query.InsertApp,
		model.NewApp(
			"admin.dev",
			"123456",
		),
	); err != nil {
		return
	}

	return
}
