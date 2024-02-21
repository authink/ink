package admin

import (
	"net/http"

	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
)

// apps godoc
//
//	@Summary		Show apps
//	@Description	Show apps
//	@Tags			app
//	@Router			/admin/apps	[get]
//	@Security		ApiKeyAuth
//	@Success		200	{array}		model.App
//	@Failure		401	{object}	inkstone.ClientError
//	@Failure		403	{object}	inkstone.ClientError
//	@Failure		500	{string}	empty
func apps(c *inkstone.Context) {
	appContext := c.App()

	apps, err := orm.App(appContext).Find()
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.JSON(http.StatusOK, apps)
}
