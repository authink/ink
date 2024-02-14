package migrate

import (
	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
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
		ink.Env.AppNameAdmin,
		"123456",
	)

	if err := orm.Transaction(ink, func(tx *sqlx.Tx) (err error) {
		if err = orm.Staff(ink).SaveWithTx(admin, tx); err != nil {
			return
		}

		return orm.App(ink).SaveWithTx(app, tx)
	}); err != nil {
		panic(err)
	}
}
