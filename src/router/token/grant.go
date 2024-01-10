package token

import (
	"net/http"
	"regexp"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/sql"
	"github.com/authink/ink.go/src/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func inkEmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return false
	}

	return matched
}

type reqGrant struct {
	AppId     int    `json:"appId" binding:"required,min=1"`
	AppSecret string `json:"appSecret" binding:"required,min=1"`
	Email     string `json:"email" binding:"required,inkEmail"`
	Password  string `json:"password" binding:"required,min=6"`
}

type resGrant struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func grant(c *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("inkEmail", inkEmailValidation)
	}

	req := &reqGrant{}
	if err := c.BindJSON(req); err != nil {
		return
	}

	extCtx := (*ext.Context)(c)
	ink := c.MustGet("ink").(*core.Ink)

	app := &model.App{}

	if err := ink.DB.Get(
		app,
		sql.Query.GetApp,
		req.AppId,
	); err != nil || !app.Active || app.Secret != util.Sha256(req.AppSecret) {
		extCtx.AbortWithClientError(ext.ERR_CLI_INVALID_APP)
		return
	}

	switch app.Name {
	case "admin.dev":
		staff := &model.Staff{}

		if err := ink.DB.Get(
			staff,
			sql.Query.GetStaff,
			req.Email,
		); err != nil || !staff.Active || staff.Departure || util.CheckPassword(staff.Password, req.Password) != nil {
			extCtx.AbortWithClientError(ext.ERR_CLI_INVALID_ACCOUNT)
			return
		}

		uuid := util.GenerateUUID()
		accessToken, err := util.GenerateToken([]byte(""), app.Id, app.Name, staff.Id, staff.Email, uuid)
		if err != nil {
			extCtx.AbortWithServerError(err)
			return
		}

		refreshToken := util.GenerateUUID()
		// accessToken identified by uuid
		authToken := model.NewAuthToken(uuid, refreshToken, app.Id, staff.Id)

		if _, err = ink.DB.NamedExec(
			sql.Query.InsertAuthToken,
			authToken,
		); err != nil {
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
		extCtx.AbortWithClientError(ext.ERR_CLI_UNSUPPORTED_APP)
	}
}
