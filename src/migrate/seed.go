package migrate

import (
	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone/app"
	"github.com/jmoiron/sqlx"
)

func Seed(appCtx *app.AppContext) {
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

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		if err = orm.Staff(appCtx).InsertTx(tx, admin); err != nil {
			return
		}

		return orm.App(appCtx).InsertTx(tx, app)
	}); err != nil {
		panic(err)
	}
}
