package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/ink.go/src/web/errs"
	"github.com/authink/ink.go/src/web/middleware"
	"github.com/authink/inkstone/web"
	"github.com/gin-gonic/gin"
)

func setupDeptGroup(gAdmin *gin.RouterGroup) {
	gDepts := gAdmin.Group(authz.Departments.Name)
	gDepts.Use(middleware.Authz(authz.Departments))
	gDepts.POST("", web.HandlerAdapter(addDept))
	gDepts.POST("levels", web.HandlerAdapter(addDeptLevel))
	gDepts.POST("staffs", web.HandlerAdapter(addDeptStaff))
}

type addDeptReq struct {
	Name   string `json:"name" binding:"required,min=6" example:"A company"`
	OwerId int    `json:"owerId" binding:"required,min=100000" example:"100000"`
}

// addDept godoc
//
//	@Summary		Add a department
//	@Description	Add a department
//	@Tags			admin_department
//	@Router			/admin/departments	[post]
//	@Security		ApiKeyAuth
//	@Param			addDeptReq	body		addDeptReq	true	"request body"
//	@Success		200			{string}	empty
//	@Failure		400			{object}	web.ClientError
//	@Failure		401			{object}	web.ClientError
//	@Failure		403			{object}	web.ClientError
//	@Failure		500			{string}	empty
func addDept(c *web.Context) {
	req := new(addDeptReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	if err := orm.Dept(c.AppContext()).Insert(models.NewDept(req.Name, req.OwerId)); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
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
	req := new(addDeptLevelReq)
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
	req := new(addDeptStaffReq)
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
