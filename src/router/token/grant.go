package token

import (
	"net/http"
	"regexp"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
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

	if app, err := ink.GetApp(req.AppId); util.CheckApp(extCtx, err, app.Active, func() bool { return util.CompareSecrets(app.Secret, req.AppSecret) }) {
		switch app.Name {
		case "admin.dev":
			staff, err := ink.GetStaffByEmail(req.Email)

			if ok := util.CheckStaff(extCtx, err, staff.Active, staff.Departure, func() bool { return util.CheckPassword(staff.Password, req.Password) == nil }); !ok {
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
