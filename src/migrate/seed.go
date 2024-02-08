package migrate

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/jmoiron/sqlx"
)

func Seed(ink *core.Ink) {
	admin := model.NewStaff(
		"admin@huoyijie.cn",
		"123456",
		"18222222222",
		true,
	)

	app := model.NewApp(
		core.APP_ADMIN_DEV,
		"123456",
	)

	if err := ink.Transaction(func(tx *sqlx.Tx) (err error) {
		if err = ink.SaveStaffWithTx(admin, tx); err != nil {
			return
		}

		return ink.SaveAppWithTx(app, tx)
	}); err != nil {
		panic(err)
	}
}
