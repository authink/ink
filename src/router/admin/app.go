package admin

import (
	"net/http"
	"time"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/authink/inkstone"
	"github.com/jmoiron/sqlx"
)

type appRes struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
}

// apps godoc
//
//	@Summary		Show apps
//	@Description	Show apps
//	@Tags			app
//	@Router			/admin/apps	[get]
//	@Security		ApiKeyAuth
//	@Success		200	{array}		appRes
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func apps(c *inkstone.Context) {
	appContext := c.App()

	apps, err := orm.App(appContext).Find()
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res []appRes
	for i := range apps {
		res = append(res, appRes{
			int(apps[i].Id),
			apps[i].CreatedAt,
			apps[i].UpdatedAt,
			apps[i].Name,
			apps[i].Active,
		})
	}

	c.JSON(http.StatusOK, res)
}

type addAppReq struct {
	Name string `json:"name" binding:"required,min=6"`
}

type addAppRes struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

// addApp godoc
//
//	@Summary		Add a app
//	@Description	Add a app
//	@Tags			app
//	@Router			/admin/apps	[post]
//	@Security		ApiKeyAuth
//	@Param			addAppReq	body		addAppReq	true	"request body"
//	@Success		200			{object}	addAppRes
//	@Failure		401			{object}	inkstone.ClientError
//	@Failure		403			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func addApp(c *inkstone.Context) {
	req := &addAppReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	secret := util.RandString(6)
	app := model.NewApp(req.Name, secret)
	if err := orm.App(c.App()).Save(app); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.JSON(http.StatusOK, &addAppRes{
		int(app.Id),
		app.Name,
		secret,
	})
}

type resetAppReq struct {
	Id int `uri:"id" binding:"required,min=100000"`
}

// resetApp godoc
//
//	@Summary		Reset a app
//	@Description	Reset a app
//	@Tags			app
//	@Router			/admin/apps/{id}/reset	[put]
//	@Security		ApiKeyAuth
//	@Param			id	path		int	true	"app id"
//	@Success		200	{object}	addAppRes
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func resetApp(c *inkstone.Context) {
	req := &resetAppReq{}
	if err := c.ShouldBindUri(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	appContext := c.App()

	var app *model.App
	secret := util.RandString(6)

	if err := inkstone.Transaction(appContext, func(tx *sqlx.Tx) (err error) {
		app, err = orm.App(appContext).Get(req.Id)
		if err != nil {
			return
		}

		app.Reset(secret)

		return orm.App(appContext).SaveWithTx(app, tx)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.JSON(http.StatusOK, &addAppRes{
		int(app.Id),
		app.Name,
		secret,
	})
}

type toggleAppReq struct {
	Id int `uri:"id" binding:"required,min=100000"`
}

type toggleAppRes struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// resetApp godoc
//
//	@Summary		Toggle a app
//	@Description	Toggle a app
//	@Tags			app
//	@Router			/admin/apps/{id}/toggle	[put]
//	@Security		ApiKeyAuth
//	@Param			id	path		int	true	"app id"
//	@Success		200	{object}	toggleAppRes
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func toggleApp(c *inkstone.Context) {
	req := &toggleAppReq{}
	if err := c.ShouldBindUri(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	appContext := c.App()

	var app *model.App

	if err := inkstone.Transaction(appContext, func(tx *sqlx.Tx) (err error) {
		app, err = orm.App(appContext).Get(req.Id)
		if err != nil {
			return
		}

		app.Active = !app.Active

		return orm.App(appContext).SaveWithTx(app, tx)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.JSON(http.StatusOK, &toggleAppRes{
		int(app.Id),
		app.Name,
		app.Active,
	})
}
