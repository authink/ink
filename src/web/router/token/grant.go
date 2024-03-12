package token

import (
	"net/http"

	"github.com/authink/ink/src/envs"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/ink/src/web/helper"
	"github.com/authink/stone/util"
	"github.com/authink/stone/web"
)

type GrantReq struct {
	AppId     int    `json:"appId" binding:"required,min=1" example:"100000"`
	AppSecret string `json:"appSecret" binding:"required,min=1" example:"123456"`
	Email     string `json:"email" binding:"required,inkEmail" example:"admin@huoyijie.cn"`
	Password  string `json:"password" binding:"required,min=6" example:"123456"`
}

type GrantRes struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    int    `json:"expires_in"`
}

// grant godoc
//
//	@Summary		Generate token
//	@Description	Generate token
//	@Tags			token
//	@Router			/token/grant [post]
//	@Param			locale		query		string		false	"locale(en, zh-CN)"
//	@Param			grantReq	body		GrantReq	true	"request body"
//	@Success		200			{object}	GrantRes
//	@Failure		400			{object}	web.ClientError
//	@Failure		500			{string}	empty
func grant(c *web.Context) {
	req := &GrantReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	var (
		appCtx = c.AppContext()
		app    models.App
	)
	app.Id = uint32(req.AppId)

	if err := orm.App(appCtx).Get(&app); helper.CheckApp(c, err, app.Active, func() bool { return helper.CompareSecrets(app.Secret, req.AppSecret) }, http.StatusBadRequest) {
		switch app.Name {
		case envs.AppNameAdmin():
			var staff = models.Staff{Email: req.Email}
			err = orm.Staff(appCtx).GetByEmail(&staff)

			if ok := helper.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return util.CheckPassword(staff.Password, req.Password) == nil }, http.StatusBadRequest); !ok {
				return
			}

			if res := generateAuthToken(c, &app, &staff); res != nil {
				c.Response(res)
			}

		default:
			c.AbortWithClientError(errs.ERR_UNSUPPORTED_APP)
		}
	}
}
