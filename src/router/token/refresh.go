package token

import (
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
)

type refreshReq struct {
	AccessToken  string `json:"access_token" binding:"required,min=1"`
	RefreshToken string `json:"refresh_token" binding:"required,min=1"`
}

func refresh(c *ext.Context) {
	req := &refreshReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(ext.ERR_BAD_REQUEST)
		return
	}

	ink := c.MustGet("ink").(*core.Ink)

	authToken, ok := checkRefreshToken(c, ink, req.RefreshToken)
	if !ok {
		return
	}

	jwtClaims, ok := util.CheckAccessToken(c, ink.Env.SecretKey, req.AccessToken, authToken.AccessToken)
	if !ok {
		return
	}

	if app, err := orm.App(ink).Get(jwtClaims.AppId); util.CheckApp(c, err, app.Active, func() bool { return true }, http.StatusBadRequest) {
		switch app.Name {
		case ink.Env.AppNameAdmin:
			staff, err := orm.Staff(ink).Get(jwtClaims.AccountId)

			if ok := util.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return true }, http.StatusBadRequest); !ok {
				return
			}

			if res := generateAuthToken(c, ink, app, staff); res != nil {
				c.JSON(http.StatusOK, res)
			}

		default:
			c.AbortWithClientError(ext.ERR_UNSUPPORTED_APP)
		}
	}
}
