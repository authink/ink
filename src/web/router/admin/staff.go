package admin

import (
	"errors"

	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/utils"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/ink/src/web/middleware"
	"github.com/authink/orm/model"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func setupStaffGroup(gAdmin *gin.RouterGroup) {
	gStaffs := gAdmin.Group(authz.Staffs.Name)
	gStaffs.Use(
		middleware.Authz(authz.Staffs),
		middleware.Log(authz.Staffs),
	)
	gStaffs.GET("", web.HandlerAdapter(staffs))
	gStaffs.POST("", web.HandlerAdapter(addStaff))
	gStaffs.PUT("", web.HandlerAdapter(updateStaff))
	gStaffs.GET("select", web.HandlerAdapter(selectStaffs))
}

type staffRes struct {
	web.Response
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
//	@Success		200		{object}	web.PagingResponse[staffRes]
//	@Failure		400		{object}	web.ClientError
//	@Failure		401		{object}	web.ClientError
//	@Failure		403		{object}	web.ClientError
//	@Failure		500		{string}	empty
func staffs(c *web.Context) {
	req := &web.PagingRequest{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	var (
		appCtx = c.AppContext()
		total  int
		staffs []*models.Staff
	)

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		if total, err = orm.Staff(appCtx).CountTx(tx); err != nil {
			return
		}

		page := model.Page{
			Offset: req.Offset,
			Limit:  req.Limit,
		}

		staffs, err = orm.Staff(appCtx).PaginationTx(tx, &page)
		return
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res = []staffRes{}
	for i := range staffs {
		staff := staffs[i]
		res = append(res, staffRes{
			Response: web.Response{
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

	c.Response(&web.PagingResponse[staffRes]{
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  total,
		Items:  res,
	})
}

type addStaffReq struct {
	Email string `json:"email" binding:"required,inkEmail" example:"example@huoyijie.cn"`
	Phone string `json:"phone" binding:"required,inkPhone" example:"18555201314"`
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
//	@Failure		400			{object}	web.ClientError
//	@Failure		401			{object}	web.ClientError
//	@Failure		403			{object}	web.ClientError
//	@Failure		500			{string}	empty
func addStaff(c *web.Context) {
	req := &addStaffReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	password := utils.RandString(6)
	staff := models.NewStaff(req.Email, password, req.Phone, false)
	if err := orm.Staff(c.AppContext()).Insert(staff); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&staffRes{
		Response: web.Response{
			Id: int(staff.Id),
		},
		Email:     staff.Email,
		Password:  password,
		Super:     staff.Super,
		Active:    staff.Active,
		Departure: staff.Departure,
	})
}

type updateStaffReq struct {
	Id              int    `json:"id" binding:"required,min=100000" example:"100000"`
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
//	@Router			/admin/staffs	[put]
//	@Security		ApiKeyAuth
//	@Param			updateStaffReq	body		updateStaffReq	true	"request body"
//	@Success		200				{object}	staffRes
//	@Failure		400				{object}	web.ClientError
//	@Failure		401				{object}	web.ClientError
//	@Failure		403				{object}	web.ClientError
//	@Failure		500				{string}	empty
func updateStaff(c *web.Context) {
	req := &updateStaffReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	var (
		appCtx   = c.AppContext()
		staff    models.Staff
		password string
	)
	staff.Id = uint32(req.Id)

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		err = orm.Staff(appCtx).GetTx(tx, &staff)
		if err != nil {
			return
		}

		if req.Phone == staff.Phone {
			return errors.New("staff's phone not changed")
		} else if req.Phone != "" {
			staff.Phone = req.Phone
		}
		if req.ActiveToggle {
			staff.Active = !staff.Active
		}
		if req.DepartureToggle {
			staff.Departure = !staff.Departure
		}
		if req.ResetPassword {
			password = utils.RandString(6)
			staff.Reset(password)
		}

		return orm.Staff(appCtx).UpdateTx(tx, &staff)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&staffRes{
		Response: web.Response{
			Id: int(staff.Id),
		},
		Email:     staff.Email,
		Phone:     staff.Phone,
		Super:     staff.Super,
		Active:    staff.Active,
		Departure: staff.Departure,
	})
}

type selectStaffRes struct {
	Id    int    `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

// selectStaffs godoc
//
//	@Summary		Select staffs
//	@Description	Select staffs
//	@Tags			admin_staff
//	@Router			/admin/staffs/select	[get]
//	@Security		ApiKeyAuth
//	@Success		200	{object}	[]selectStaffRes
//	@Failure		400	{object}	web.ClientError
//	@Failure		401	{object}	web.ClientError
//	@Failure		403	{object}	web.ClientError
//	@Failure		500	{string}	empty
func selectStaffs(c *web.Context) {
	staffs, err := orm.Staff(c.AppContext()).SelectStaffs()
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res = []selectStaffRes{}
	for i := range staffs {
		res = append(res, selectStaffRes{
			Id:    int(staffs[i].Id),
			Email: staffs[i].Email,
		})
	}

	c.Response(res)
}
