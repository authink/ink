package admin

import (
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type appRes struct {
	inkstone.Response
	Name   string `json:"name,omitempty"`
	Secret string `json:"secret,omitempty"`
	Active bool   `json:"active"`
}

// apps godoc
//
//	@Summary		Show apps
//	@Description	Show apps
//	@Tags			admin_app
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
			Response: inkstone.Response{
				Id:        int(apps[i].Id),
				CreatedAt: apps[i].CreatedAt,
				UpdatedAt: apps[i].UpdatedAt,
			},
			Name:   apps[i].Name,
			Active: apps[i].Active,
		})
	}

	c.Response(res)
}

type addAppReq struct {
	Name string `json:"name" binding:"required,min=6" example:"appmock"`
}

// addApp godoc
//
//	@Summary		Add a app
//	@Description	Add a app
//	@Tags			admin_app
//	@Router			/admin/apps	[post]
//	@Security		ApiKeyAuth
//	@Param			addAppReq	body		addAppReq	true	"request body"
//	@Success		200			{object}	appRes
//	@Failure		400			{object}	inkstone.ClientError
//	@Failure		401			{object}	inkstone.ClientError
//	@Failure		403			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func addApp(c *inkstone.Context) {
	req := new(addAppReq)
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

	c.Response(&appRes{
		Response: inkstone.Response{
			Id: int(app.Id),
		},
		Name:   app.Name,
		Secret: secret,
		Active: app.Active,
	})
}

type updateAppParam struct {
	Id int `uri:"id" binding:"required,min=100000"`
}

type updateAppReq struct {
	ResetSecret  bool `json:"resetSecret" example:"false"`
	ActiveToggle bool `json:"activeToggle" example:"true"`
}

// updateApp godoc
//
//	@Summary		Update a app
//	@Description	Update a app
//	@Tags			admin_app
//	@Router			/admin/apps/{id}	[put]
//	@Security		ApiKeyAuth
//	@Param			id				path		int				true	"app id"
//	@Param			updateAppReq	body		updateAppReq	true	"request body"
//	@Success		200				{object}	appRes
//	@Failure		400				{object}	inkstone.ClientError
//	@Failure		401				{object}	inkstone.ClientError
//	@Failure		403				{object}	inkstone.ClientError
//	@Failure		500				{string}	empty
func updateApp(c *inkstone.Context) {
	param := new(updateAppParam)

	if err := c.ShouldBindUri(param); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	req := new(updateAppReq)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(inkstone.ValidationNotAllFieldsZero, req)
	}

	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var (
		appContext = c.App()
		app        *model.App
		secret     string
	)

	if err := inkstone.Transaction(appContext, func(tx *sqlx.Tx) (err error) {
		app, err = orm.App(appContext).GetWithTx(param.Id, tx)
		if err != nil {
			return
		}

		if req.ResetSecret {
			secret = util.RandString(6)
			app.Reset(secret)
		}
		if req.ActiveToggle {
			app.Active = !app.Active
		}

		return orm.App(appContext).SaveWithTx(app, tx)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&appRes{
		Response: inkstone.Response{
			Id: int(app.Id),
		},
		Name:   app.Name,
		Secret: secret,
		Active: app.Active,
	})
}
