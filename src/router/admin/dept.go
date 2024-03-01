package admin

import (
	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
)

func setupDeptGroup(gAdmin *gin.RouterGroup) {
	gDepts := gAdmin.Group(authz.Departments.Name)
	gDepts.Use(middleware.Authz(authz.Departments))
	gDepts.POST("", inkstone.HandlerAdapter(addDept))
	gDepts.POST("levels", inkstone.HandlerAdapter(addDeptLevel))
	gDepts.POST("staffs", inkstone.HandlerAdapter(addDeptStaff))
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
//	@Failure		400			{object}	inkstone.ClientError
//	@Failure		401			{object}	inkstone.ClientError
//	@Failure		403			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func addDept(c *inkstone.Context) {
	req := new(addDeptReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	if err := orm.Dept(c.AppContext()).Insert(model.NewDept(req.Name, req.OwerId)); err != nil {
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
//	@Failure		400				{object}	inkstone.ClientError
//	@Failure		401				{object}	inkstone.ClientError
//	@Failure		403				{object}	inkstone.ClientError
//	@Failure		500				{string}	empty
func addDeptLevel(c *inkstone.Context) {
	req := new(addDeptLevelReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	if err := orm.DeptLevel(c.AppContext()).Insert(&model.DeptLevel{
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
//	@Failure		400				{object}	inkstone.ClientError
//	@Failure		401				{object}	inkstone.ClientError
//	@Failure		403				{object}	inkstone.ClientError
//	@Failure		500				{string}	empty
func addDeptStaff(c *inkstone.Context) {
	req := new(addDeptStaffReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	if err := orm.DeptStaff(c.AppContext()).Insert(&model.DeptStaff{
		DeptId:  uint32(req.DeptId),
		StaffId: uint32(req.StaffId),
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Empty()
}
