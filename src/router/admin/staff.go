package admin

import (
	"time"

	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type staffRes struct {
	Id        int        `json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	Super     bool       `json:"super"`
	Active    bool       `json:"active"`
	Departure bool       `json:"departure"`
}

// staffs godoc
//
//	@Summary		Show staffs
//	@Description	Show staffs
//	@Tags			admin_staff
//	@Router			/admin/staffs	[get]
//	@Security		ApiKeyAuth
//	@Param			offset	query		int	false	"offset"
//	@Param			limit	query		int	true	"limit"
//	@Success		200		{object}	inkstone.PagingResponse[staffRes]
//	@Failure		400		{object}	inkstone.ClientError
//	@Failure		401		{object}	inkstone.ClientError
//	@Failure		403		{object}	inkstone.ClientError
//	@Failure		500		{string}	empty
func staffs(c *inkstone.Context) {
	appContext := c.App()

	req := new(inkstone.PagingRequest)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var total int
	var staffs []model.Staff

	if err := inkstone.Transaction(appContext, func(tx *sqlx.Tx) (err error) {
		if total, err = orm.Staff(appContext).CountWithTx(tx); err != nil {
			return
		}

		staffs, err = orm.Staff(appContext).PaginationWithTx(req.Offset, req.Limit, tx)
		return
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res []staffRes
	for i := range staffs {
		staff := &staffs[i]
		res = append(res, staffRes{
			int(staff.Id),
			&staff.CreatedAt,
			&staff.UpdatedAt,
			staff.Email,
			staff.Phone,
			staff.Super,
			staff.Active,
			staff.Departure,
		})
	}

	c.Response(&inkstone.PagingResponse[staffRes]{
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  total,
		Items:  res,
	})
}

type addStaffReq struct {
	Email string `json:"email" binding:"required,inkEmail" example:"example@huoyijie.cn"`
	Phone string `json:"phone" binding:"required,min=11,max=11" example:"18555201314"`
	Super bool   `json:"super" example:"false"`
}

type addStaffRes struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// addStaff godoc
//
//	@Summary		Add a staff
//	@Description	Add a staff
//	@Tags			admin_staff
//	@Router			/admin/staffs	[post]
//	@Security		ApiKeyAuth
//	@Param			addStaffReq	body		addStaffReq	true	"request body"
//	@Success		200			{object}	addStaffRes
//	@Failure		400			{object}	inkstone.ClientError
//	@Failure		401			{object}	inkstone.ClientError
//	@Failure		403			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func addStaff(c *inkstone.Context) {
	req := new(addStaffReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	password := util.RandString(6)
	staff := model.NewStaff(req.Email, password, req.Phone, req.Super)
	if err := orm.Staff(c.App()).Save(staff); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&addStaffRes{
		Id:       int(staff.Id),
		Email:    staff.Email,
		Password: password,
	})
}

type updateStaffParam struct {
	Id int `uri:"id" binding:"required,min=100000"`
}

type updateStaffReq struct {
	Phone           string `json:"phone" binding:"omitempty,min=11,max=11" example:"18555201314"`
	SuperToggle     bool   `json:"superToggle" example:"false"`
	ActiveToggle    bool   `json:"activeToggle" example:"true"`
	DepartureToggle bool   `json:"departureToggle" example:"false"`
	ResetPassword   bool   `json:"resetPassword" example:"false"`
}

// updateStaff godoc
//
//	@Summary		Update a staff
//	@Description	Update a staff
//	@Tags			admin_staff
//	@Router			/admin/staffs/{id}	[put]
//	@Security		ApiKeyAuth
//	@Param			id				path		int				true	"staff id"
//	@Param			updateStaffReq	body		updateStaffReq	true	"request body"
//	@Success		200				{object}	staffRes
//	@Failure		400				{object}	inkstone.ClientError
//	@Failure		401				{object}	inkstone.ClientError
//	@Failure		403				{object}	inkstone.ClientError
//	@Failure		500				{string}	empty
func updateStaff(c *inkstone.Context) {
	param := new(updateStaffParam)

	if err := c.ShouldBindUri(param); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	req := new(updateStaffReq)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(inkstone.ValidationNotAllFieldsZero, req)
	}

	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	appContext := c.App()

	var staff *model.Staff

	if err := inkstone.Transaction(appContext, func(tx *sqlx.Tx) (err error) {
		staff, err = orm.Staff(appContext).GetWithTx(param.Id, tx)
		if err != nil {
			return
		}

		if req.Phone != "" {
			staff.Phone = req.Phone
		}
		if req.SuperToggle {
			staff.Super = !staff.Super
		}
		if req.ActiveToggle {
			staff.Active = !staff.Active
		}
		if req.DepartureToggle {
			staff.Departure = !staff.Departure
		}
		if req.ResetPassword {
			staff.Reset(util.RandString(6))
		}

		return orm.Staff(appContext).SaveWithTx(staff, tx)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&staffRes{
		Id:        int(staff.Id),
		Email:     staff.Email,
		Phone:     staff.Phone,
		Super:     staff.Super,
		Active:    staff.Active,
		Departure: staff.Departure,
	})
}
