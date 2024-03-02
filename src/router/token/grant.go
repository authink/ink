package token

import (
	"net/http"

	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	u "github.com/authink/inkstone/util"
	"github.com/authink/inkstone/web"
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
//	@Param			lang		query		string		false	"language"
//	@Param			grantReq	body		GrantReq	true	"request body"
//	@Success		200			{object}	GrantRes
//	@Failure		400			{object}	web.ClientError
//	@Failure		500			{string}	empty
func grant(c *web.Context) {
	req := new(GrantReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	appCtx := c.AppContext()

	if app, err := orm.App(appCtx).Get(req.AppId); util.CheckApp(c, err, app.Active, func() bool { return util.CompareSecrets(app.Secret, req.AppSecret) }, http.StatusBadRequest) {
		switch app.Name {
		case env.AppNameAdmin():
			staff, err := orm.Staff(appCtx).GetByEmail(req.Email)

			if ok := util.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return u.CheckPassword(staff.Password, req.Password) == nil }, http.StatusBadRequest); !ok {
				return
			}

			if res := generateAuthToken(c, app, staff); res != nil {
				c.Response(res)
			}

		default:
			c.AbortWithClientError(errors.ERR_UNSUPPORTED_APP)
		}
	}
}
