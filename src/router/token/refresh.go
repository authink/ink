package token

import (
	"net/http"

	"github.com/authink/ink.go/src/env"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/authink/inkstone"
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
//	@Failure		400			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func refresh(c *inkstone.Context) {
	req := &refreshReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	authToken, ok := checkRefreshToken(c, req.RefreshToken)
	if !ok {
		return
	}

	appContext := c.App()

	jwtClaims, ok := util.CheckAccessToken(c, appContext.SecretKey, req.AccessToken, authToken.AccessToken)
	if !ok {
		return
	}

	if app, err := orm.App(appContext).Get(jwtClaims.AppId); util.CheckApp(c, err, app.Active, func() bool { return true }, http.StatusBadRequest) {
		switch app.Name {
		case env.AppNameAdmin():
			staff, err := orm.Staff(appContext).Get(jwtClaims.AccountId)

			if ok := util.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return true }, http.StatusBadRequest); !ok {
				return
			}

			if res := generateAuthToken(c, app, staff); res != nil {
				c.JSON(http.StatusOK, res)
			}

		default:
			c.AbortWithClientError(errors.ERR_UNSUPPORTED_APP)
		}
	}
}
