package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/util"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func setupStaffGroup(gAdmin *gin.RouterGroup) {
	gStaffs := gAdmin.Group(authz.Staffs.Name)
	gStaffs.Use(middleware.Authz(authz.Staffs))
	gStaffs.GET("", inkstone.HandlerAdapter(staffs))
	gStaffs.POST("", inkstone.HandlerAdapter(addStaff))
	gStaffs.PUT(":id", inkstone.HandlerAdapter(updateStaff))
}

type staffRes struct {
	inkstone.Response
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Super     bool   `json:"super"`
	Active    bool   `json:"active"`
	Departure bool   `json:"departure"`
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
	appCtx := c.AppContext()

	req := new(inkstone.PagingRequest)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var total int
	var staffs []model.Staff

	if err := inkstone.Transaction(appCtx, func(tx *sqlx.Tx) (err error) {
		if total, err = orm.Staff(appCtx).CountWithTx(tx); err != nil {
			return
		}

		staffs, err = orm.Staff(appCtx).PaginationWithTx(req.Offset, req.Limit, tx)
		return
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res []staffRes
	for i := range staffs {
		staff := &staffs[i]
		res = append(res, staffRes{
			Response: inkstone.Response{
				Id:        int(staff.Id),
				CreatedAt: staff.CreatedAt,
				UpdatedAt: staff.UpdatedAt,
			},
			Email:     staff.Email,
			Phone:     staff.Phone,
			Super:     staff.Super,
			Active:    staff.Active,
			Departure: staff.Departure,
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
}

// addStaff godoc
//
//	@Summary		Add a staff
//	@Description	Add a staff
//	@Tags			admin_staff
//	@Router			/admin/staffs	[post]
//	@Security		ApiKeyAuth
//	@Param			addStaffReq	body		addStaffReq	true	"request body"
//	@Success		200			{object}	staffRes
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
	staff := model.NewStaff(req.Email, password, req.Phone, false)
	if err := orm.Staff(c.AppContext()).Insert(staff); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&staffRes{
		Response: inkstone.Response{
			Id: int(staff.Id),
		},
		Email:     staff.Email,
		Password:  password,
		Super:     staff.Super,
		Active:    staff.Active,
		Departure: staff.Departure,
	})
}

type updateStaffParam struct {
	Id int `uri:"id" binding:"required,min=100000"`
}

type updateStaffReq struct {
	Phone           string `json:"phone" binding:"omitempty,min=11,max=11" example:"18555201314"`
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

	var (
		appCtx   = c.AppContext()
		staff    *model.Staff
		password string
	)

	if err := inkstone.Transaction(appCtx, func(tx *sqlx.Tx) (err error) {
		staff, err = orm.Staff(appCtx).GetWithTx(param.Id, tx)
		if err != nil {
			return
		}

		if req.Phone != "" {
			staff.Phone = req.Phone
		}
		if req.ActiveToggle {
			staff.Active = !staff.Active
		}
		if req.DepartureToggle {
			staff.Departure = !staff.Departure
		}
		if req.ResetPassword {
			password = util.RandString(6)
			staff.Reset(password)
		}

		return orm.Staff(appCtx).SaveWithTx(staff, tx)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&staffRes{
		Response: inkstone.Response{
			Id: int(staff.Id),
		},
		Email:     staff.Email,
		Phone:     staff.Phone,
		Super:     staff.Super,
		Active:    staff.Active,
		Departure: staff.Departure,
	})
}
