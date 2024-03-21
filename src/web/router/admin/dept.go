package admin

import (
	"strings"

	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/ink/src/web/middleware"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func setupDeptGroup(gAdmin *gin.RouterGroup) {
	gDepts := gAdmin.Group(authz.Departments.Name)
	gDepts.Use(
		middleware.Authz(authz.Departments),
		middleware.Log(authz.Departments),
	)
	gDepts.GET("", web.HandlerAdapter(depts))
	gDepts.GET(":name/unique", web.HandlerAdapter(unique))
	gDepts.POST("", web.HandlerAdapter(saveDept))
	gDepts.POST("levels", web.HandlerAdapter(addDeptLevel))
	gDepts.POST("staffs", web.HandlerAdapter(addDeptStaff))
}

type deptRes struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	OwnerId   int    `json:"ownerId"`
	OwnerName string `json:"ownerName"`
}

// depts godoc
//
//	@Summary		Show depts
//	@Description	Show depts
//	@Tags			admin_department
//	@Router			/admin/departments	[get]
//	@Security		ApiKeyAuth
//	@Success		200	{array}		deptRes
//	@Failure		401	{object}	web.ClientError
//	@Failure		403	{object}	web.ClientError
//	@Failure		500	{string}	empty
func depts(c *web.Context) {
	depts, err := orm.Dept(c.AppContext()).Find()
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res = []deptRes{}
	for _, v := range depts {
		res = append(res, deptRes{
			Id:        int(v.Id),
			Name:      v.Name,
			Active:    v.Active,
			OwnerId:   int(v.OwnerId),
			OwnerName: strings.Split(v.OwnerEmail, "@")[0],
		})
	}

	c.Response(res)
}

type uniqueReq struct {
	Name string `uri:"name" binding:"required,min=2"`
}

// unique godoc
//
//	@Summary		Check if unique name
//	@Description	Check if unique name
//	@Tags			admin_department
//	@Router			/admin/departments/{name}/unique	[get]
//	@Security		ApiKeyAuth
//	@Param			name	path		string	true	"dept name"
//	@Success		200		{boolean}	boolean
//	@Failure		400		{object}	web.ClientError
//	@Failure		401		{object}	web.ClientError
//	@Failure		403		{object}	web.ClientError
//	@Failure		500		{string}	empty
func unique(c *web.Context) {
	req := &uniqueReq{}
	err := c.ShouldBindUri(req)
	if err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	uni, err := orm.Dept(c.AppContext()).Unique(&models.CheckUnique{
		Name: req.Name,
	})
	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(uni)
}

type saveDeptReq struct {
	Id       uint32 `json:"id" example:"100000"`
	Name     string `json:"name" binding:"required,min=2" example:"New Department"`
	OwnerId  int    `json:"ownerId" binding:"required,min=100000" example:"100000"`
	ParentId int    `json:"parentId" binding:"required,min=100000" example:"100000"`
}

// saveDept godoc
//
//	@Summary		Save a department
//	@Description	Save a department
//	@Tags			admin_department
//	@Router			/admin/departments	[post]
//	@Security		ApiKeyAuth
//	@Param			saveDeptReq	body		saveDeptReq	true	"request body"
//	@Success		200			{int}		int
//	@Failure		400			{object}	web.ClientError
//	@Failure		401			{object}	web.ClientError
//	@Failure		403			{object}	web.ClientError
//	@Failure		500			{string}	empty
func saveDept(c *web.Context) {
	req := &saveDeptReq{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	dept := &models.Department{
		Name:    req.Name,
		OwnerId: uint32(req.OwnerId),
	}
	dept.Id = req.Id
	appCtx := c.AppContext()

	if dept.Id == 0 {
		err = appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
			if err = orm.Dept(appCtx).InsertTx(tx, dept); err != nil {
				return
			}

			if req.ParentId > 0 {
				err = orm.DeptLevel(appCtx).InsertTx(tx, &models.DeptLevel{
					DeptId:    uint32(req.ParentId),
					SubDeptId: dept.Id,
				})
			}
			return
		})
	} else {
		err = orm.Dept(appCtx).Update(dept)
	}

	if err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(dept.Id)
}

type addDeptLevelReq struct {
	DeptId    int `json:"deptId" binding:"required,min=100000" example:"100000"`
	SubDeptId int `json:"subDeptId" binding:"required,min=100000" example:"100001"`
}

// addDeptLevel godoc
//
//	@Summary		Add a department level
//	@Description	Add a department level
//	@Tags			admin_department
//	@Router			/admin/departments/levels	[post]
//	@Security		ApiKeyAuth
//	@Param			addDeptLevelReq	body		addDeptLevelReq	true	"request body"
//	@Success		200				{string}	empty
//	@Failure		400				{object}	web.ClientError
//	@Failure		401				{object}	web.ClientError
//	@Failure		403				{object}	web.ClientError
//	@Failure		500				{string}	empty
func addDeptLevel(c *web.Context) {
	req := &addDeptLevelReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	if err := orm.DeptLevel(c.AppContext()).Insert(&models.DeptLevel{
		DeptId:    uint32(req.DeptId),
		SubDeptId: uint32(req.SubDeptId),
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}

type addDeptStaffReq struct {
	DeptId  int `json:"deptId" binding:"required,min=100000" example:"100000"`
	StaffId int `json:"staffId" binding:"required,min=100000" example:"100000"`
}

// addDeptStaff godoc
//
//	@Summary		Add a staff to the department
//	@Description	Add a staff to the department
//	@Tags			admin_department
//	@Router			/admin/departments/staffs	[post]
//	@Security		ApiKeyAuth
//	@Param			addDeptStaffReq	body		addDeptStaffReq	true	"request body"
//	@Success		200				{string}	empty
//	@Failure		400				{object}	web.ClientError
//	@Failure		401				{object}	web.ClientError
//	@Failure		403				{object}	web.ClientError
//	@Failure		500				{string}	empty
func addDeptStaff(c *web.Context) {
	req := &addDeptStaffReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	if err := orm.DeptStaff(c.AppContext()).Insert(&models.DeptStaff{
		DeptId:  uint32(req.DeptId),
		StaffId: uint32(req.StaffId),
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}
