package migrate

import (
	"github.com/authink/ink/src/envs"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/inkstone/app"
	"github.com/jmoiron/sqlx"
)

func Seed(appCtx *app.AppContext) {
	admin := models.NewStaff(
		"admin@huoyijie.cn",
		"123456",
		"18222222222",
		true,
	)

	app := models.NewApp(
		envs.AppNameAdmin(),
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
