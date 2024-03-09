package admin

import (
	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/utils"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/ink/src/web/middleware"
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
		app := apps[i]
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
	req := &addAppReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	secret := utils.RandString(6)
	app := models.NewApp(req.Name, secret)
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
	param := &updateAppParam{}

	if err := c.ShouldBindUri(param); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	req := &updateAppReq{}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(web.ValidationNotAllFieldsZero, req)
	}

	var currentApp = c.MustGet("app").(*models.App)
	if err := c.ShouldBindJSON(req); err != nil || (req.ActiveToggle && param.Id == int(currentApp.Id)) {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	var (
		appCtx = c.AppContext()
		app    models.App
		secret string
	)
	app.Id = uint32(param.Id)

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		err = orm.App(appCtx).GetTx(tx, &app)
		if err != nil {
			return
		}

		if req.ResetSecret {
			secret = utils.RandString(6)
			app.Reset(secret)
		}
		if req.ActiveToggle {
			app.Active = !app.Active
		}

		return orm.App(appCtx).UpdateTx(tx, &app)
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
