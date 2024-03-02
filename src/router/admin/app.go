package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func setupAppGroup(gAdmin *gin.RouterGroup) {
	gApps := gAdmin.Group(authz.Apps.Name)
	gApps.Use(middleware.Authz(authz.Apps))
	gApps.GET("", web.HandlerAdapter(apps))
	gApps.POST("", web.HandlerAdapter(addApp))
	gApps.PUT(":id", web.HandlerAdapter(updateApp))
}

type appRes struct {
	web.Response
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
//	@Failure		401	{object}	web.ClientError
//	@Failure		403	{object}	web.ClientError
//	@Failure		500	{string}	empty
func apps(c *web.Context) {
	appCtx := c.AppContext()

	apps, err := orm.App(appCtx).Find()
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res = []appRes{}
	for i := range apps {
		app := &apps[i]
		res = append(res, appRes{
			Response: web.Response{
				Id:        int(app.Id),
				CreatedAt: app.CreatedAt,
				UpdatedAt: app.UpdatedAt,
			},
			Name:   app.Name,
			Active: app.Active,
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
//	@Failure		400			{object}	web.ClientError
//	@Failure		401			{object}	web.ClientError
//	@Failure		403			{object}	web.ClientError
//	@Failure		500			{string}	empty
func addApp(c *web.Context) {
	req := new(addAppReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	secret := util.RandString(6)
	app := model.NewApp(req.Name, secret)
	if err := orm.App(c.AppContext()).Insert(app); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&appRes{
		Response: web.Response{
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
//	@Failure		400				{object}	web.ClientError
//	@Failure		401				{object}	web.ClientError
//	@Failure		403				{object}	web.ClientError
//	@Failure		500				{string}	empty
func updateApp(c *web.Context) {
	param := new(updateAppParam)

	if err := c.ShouldBindUri(param); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	req := new(updateAppReq)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(web.ValidationNotAllFieldsZero, req)
	}

	var currentApp = c.MustGet("app").(*model.App)
	if err := c.ShouldBindJSON(req); err != nil || (req.ActiveToggle && param.Id == int(currentApp.Id)) {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var (
		appCtx = c.AppContext()
		app    *model.App
		secret string
	)

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		app, err = orm.App(appCtx).GetTx(tx, param.Id)
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

		return orm.App(appCtx).UpdateTx(tx, app)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&appRes{
		Response: web.Response{
			Id: int(app.Id),
		},
		Name:   app.Name,
		Secret: secret,
		Active: app.Active,
	})
}
