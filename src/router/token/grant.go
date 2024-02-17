package token

import (
	"net/http"
	"regexp"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
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

type GrantReq struct {
	AppId     int    `json:"appId" binding:"required,min=1" example:"100000"`
	AppSecret string `json:"appSecret" binding:"required,min=1" example:"123456"`
	Email     string `json:"email" binding:"required,inkEmail" example:"admin@huoyijie.cn"`
	Password  string `json:"password" binding:"required,min=6" example:"123456"`
}

type GrantRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// grant godoc
//
//	@Summary		Generate token
//	@Description	Generate token
//	@Tags			token
//	@Router			/token/grant [post]
//	@Param			grantReq	body		grantReq	true	"request body"
//	@Success		200			{object}	grantRes
//	@Failure		400			{object}	ext.ClientError
//	@Failure		500			{string}	empty
func grant(c *ext.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("inkEmail", inkEmailValidation)
	}

	req := &GrantReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(ext.ERR_BAD_REQUEST)
		return
	}

	ink := c.MustGet("ink").(*core.Ink)

	if app, err := orm.App(ink).Get(req.AppId); util.CheckApp(c, err, app.Active, func() bool { return util.CompareSecrets(app.Secret, req.AppSecret) }, http.StatusBadRequest) {
		switch app.Name {
		case ink.Env.AppNameAdmin:
			staff, err := orm.Staff(ink).GetByEmail(req.Email)

			if ok := util.CheckStaff(c, err, staff.Active, staff.Departure, func() bool { return util.CheckPassword(staff.Password, req.Password) == nil }, http.StatusBadRequest); !ok {
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
