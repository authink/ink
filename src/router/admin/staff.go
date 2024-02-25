package admin

import (
	"net/http"
	"time"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
)

type staffRes struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Super     bool      `json:"super"`
	Active    bool      `json:"active"`
	Departure bool      `json:"departure"`
}

// staffs godoc
//
//	@Summary		Show staffs
//	@Description	Show staffs
//	@Tags			admin_staff
//	@Router			/admin/staffs	[get]
//	@Security		ApiKeyAuth
//	@Param			offset	query		int	true	"offset"
//	@Param			limit	query		int	true	"limit"
//	@Success		200		{object}	pageRes[staffRes]
//	@Failure		401		{object}	inkstone.ClientError
//	@Failure		403		{object}	inkstone.ClientError
//	@Failure		500		{string}	empty
func staffs(c *inkstone.Context) {
	appContext := c.App()

	req := &pageReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	total, err := orm.Staff(appContext).Count()
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	staffs, err := orm.Staff(appContext).Pagination(req.Offset, req.Limit)
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res []staffRes
	for i := range staffs {
		staff := &staffs[i]
		res = append(res, staffRes{
			int(staff.Id),
			staff.CreatedAt,
			staff.UpdatedAt,
			staff.Email,
			staff.Phone,
			staff.Super,
			staff.Active,
			staff.Departure,
		})
	}

	c.JSON(http.StatusOK, &pageRes[staffRes]{
		total,
		req.Offset,
		req.Limit,
		res,
	})
}
