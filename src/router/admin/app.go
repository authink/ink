package admin

import (
	"net/http"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/ext"
	"github.com/authink/ink.go/src/orm"
)

// apps godoc
//
//	@Summary		Show apps
//	@Description	Show apps
//	@Tags			app
//	@Router			/admin/apps [get]
//	@Security		ApiKeyAuth
//	@Success		200	{array}		model.App
//	@Failure		401	{object}	ext.ClientError
//	@Failure		403	{object}	ext.ClientError
//	@Failure		500	{string}	empty
func apps(c *ext.Context) {
	ink := c.MustGet("ink").(*core.Ink)

	apps, err := orm.App(ink).Find()
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.JSON(http.StatusOK, apps)
}
