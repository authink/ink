package token

import (
	libsql "database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	authToken, err := ink.GetByRefreshToken(req.RefreshToken)
	if err != nil {
		if !errors.Is(err, libsql.ErrNoRows) {
			extCtx.AbortWithServerError(err)
			return
		}
		extCtx.AbortWithClientError(ext.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	_, err = ink.DeleteToken(int(authToken.Id))
	if err != nil {
		extCtx.AbortWithServerError(err)
		return
	}

	if time.Now().After(authToken.CreatedAt.Add(7 * 24 * time.Hour)) {
		extCtx.AbortWithClientError(ext.ERR_INVALID_REFRESH_TOKEN)
		return
	}

	jwtClaims, err := util.VerifyToken(
		ink.Env.SecretKey,
		req.AccessToken,
	)

	if (err != nil && !errors.Is(err, jwt.ErrTokenExpired)) || jwtClaims.ID != authToken.AccessToken {
		extCtx.AbortWithClientError(ext.ERR_INVALID_ACCESS_TOKEN)
		return
	}

	app, err := ink.GetApp(jwtClaims.AppId)
	if err != nil || !app.Active {
		if err != nil && !errors.Is(err, libsql.ErrNoRows) {
			extCtx.AbortWithServerError(err)
			return
		}
		extCtx.AbortWithClientError(ext.ERR_INVALID_APP)
		return
	}

	switch app.Name {
	case "admin.dev":
		staff, err := ink.GetStaff(jwtClaims.AccountId)

		if err != nil || !staff.Active || staff.Departure {
			if err != nil && !errors.Is(err, libsql.ErrNoRows) {
				extCtx.AbortWithServerError(err)
				return
			}
			extCtx.AbortWithClientError(ext.ERR_INVALID_ACCOUNT)
			return
		}

		// generate new token
		uuid := util.GenerateUUID()
		accessToken, err := util.GenerateToken(ink.Env.SecretKey, app.Id, app.Name, staff.Id, staff.Email, uuid)
		if err != nil {
			extCtx.AbortWithServerError(err)
			return
		}

		refreshToken := util.GenerateUUID()
		// accessToken identified by uuid
		authToken := model.NewAuthToken(uuid, refreshToken, app.Id, staff.Id)

		if _, err = ink.SaveToken(authToken); err != nil {
			extCtx.AbortWithServerError(err)
			return
		}

		res := &resGrant{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			TokenType:    "Bearer",
			ExpiresIn:    7200,
		}
		c.JSON(http.StatusOK, res)
	default:
		extCtx.AbortWithClientError(ext.ERR_UNSUPPORTED_APP)
	}
}
