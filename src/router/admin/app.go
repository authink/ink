package admin

import (
	"net/http"
	"time"

	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
)

type appRes struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
}

// apps godoc
//
//	@Summary		Show apps
//	@Description	Show apps
//	@Tags			app
//	@Router			/admin/apps	[get]
//	@Security		ApiKeyAuth
//	@Success		200	{array}		appRes
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

	var res []appRes
	for i := range apps {
		res = append(res, appRes{
			int(apps[i].Id),
			apps[i].CreatedAt,
			apps[i].UpdatedAt,
			apps[i].Name,
			apps[i].Active,
		})
	}

	c.JSON(http.StatusOK, res)
}
