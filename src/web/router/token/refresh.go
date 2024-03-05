package token

import (
	"net/http"

	"github.com/authink/ink.go/src/envs"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/web/errs"
	"github.com/authink/ink.go/src/web/helper"
	"github.com/authink/inkstone/web"
)

type refreshReq struct {
	AccessToken  string `json:"access_token" binding:"required,min=1"`
	RefreshToken string `json:"refresh_token" binding:"required,min=1"`
}

// refresh godoc
//
//	@Summary		Refresh token
//	@Description	Refresh token
//	@Tags			token
//	@Router			/token/refresh [post]
//	@Param			refreshReq	body		refreshReq	true	"request body"
//	@Success		200			{object}	GrantRes
//	@Failure		400			{object}	web.ClientError
//	@Failure		500			{string}	empty
func refresh(c *web.Context) {
	req := &refreshReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	authToken, ok := checkRefreshToken(c, req.RefreshToken)
	if !ok {
		return
	}

	appCtx := c.AppContext()

	jwtClaims, ok := helper.CheckAccessToken(c, appCtx.SecretKey, req.AccessToken, authToken.AccessToken)
	if !ok {
		return
	}

	if app, err := orm.App(appCtx).Get(jwtClaims.AppId); helper.CheckApp(c, err, app.Active, func() bool { return true }, http.StatusBadRequest) {
		switch app.Name {
		case envs.AppNameAdmin():
			staff, err := orm.Staff(appCtx).Get(jwtClaims.AccountId)

			if ok := helper.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return true }, http.StatusBadRequest); !ok {
				return
			}

			if res := generateAuthToken(c, app, staff); res != nil {
				c.Response(res)
			}

		default:
			c.AbortWithClientError(errs.ERR_UNSUPPORTED_APP)
		}
	}
}
