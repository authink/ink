package token

import (
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/util"
	"github.com/gin-gonic/gin"
)

type reqRefresh struct {
	AccessToken  string `json:"access_token" binding:"required,min=1"`
	RefreshToken string `json:"refresh_token" binding:"required,min=1"`
}

func refresh(c *gin.Context) {
	req := &reqRefresh{}
	if err := c.BindJSON(req); err != nil {
		return
	}

	extCtx := (*ext.Context)(c)
	ink := c.MustGet("ink").(*core.Ink)

	authToken, ok := checkRefreshToken(extCtx, ink, req.RefreshToken)
	if !ok {
		return
	}

	jwtClaims, ok := util.CheckAccessToken(extCtx, ink.Env.SecretKey, req.AccessToken, authToken.AccessToken)
	if !ok {
		return
	}

	if app, err := ink.GetApp(jwtClaims.AppId); util.CheckApp(extCtx, err, app.Active, func() bool { return true }) {
		switch app.Name {
		case "admin.dev":
			staff, err := ink.GetStaff(jwtClaims.AccountId)

			if ok := util.CheckStaff(extCtx, err, staff.Active, staff.Departure, func() bool { return true }); !ok {
				return
			}

			if res := generateAuthToken(extCtx, ink, app, staff); res != nil {
				c.JSON(http.StatusOK, res)
			}

		default:
			extCtx.AbortWithClientError(ext.ERR_UNSUPPORTED_APP)
		}
	}
}
