package migrate

import (
	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

func Seed(appContext *inkstone.AppContext) {
	admin := model.NewStaff(
		"admin@huoyijie.cn",
		"123456",
		"18222222222",
		true,
	)

	app := model.NewApp(
		env.AppNameAdmin(),
		"123456",
	)

	if err := inkstone.Transaction(appContext, func(tx *sqlx.Tx) (err error) {
		if err = orm.Staff(appContext).SaveWithTx(admin, tx); err != nil {
			return
		}

		return orm.App(appContext).SaveWithTx(app, tx)
	}); err != nil {
		panic(err)
	}
}
