package token

import (
	"net/http"
	"regexp"

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
	AppId     uint32 `json:"appId" binding:"required,min=1"`
	AppSecret string `json:"appSecret" binding:"required,min=1"`
	Email     string `json:"email" binding:"required,inkEmail"`
	Password  string `json:"password" binding:"required,min=6"`
}

func grant(c *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("inkEmail", inkEmailValidation)
	}

	req := &reqGrant{}
	if err := c.BindJSON(req); err != nil {
		return
	}

	c.JSON(http.StatusOK, req)
}
