package admin

import (
	errs "errors"

	"github.com/authink/ink.go/src/authz"
	"github.com/authink/ink.go/src/errors"
	"github.com/authink/ink.go/src/middleware"
	"github.com/authink/ink.go/src/model"
	"github.com/authink/ink.go/src/orm"
	"github.com/authink/inkstone"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func setupGroupGroup(gAdmin *gin.RouterGroup) {
	gGroups := gAdmin.Group(authz.Groups.Name)
	gGroups.Use(middleware.Authz(authz.Groups))
	gGroups.GET("", inkstone.HandlerAdapter(groups))
	gGroups.POST("", inkstone.HandlerAdapter(addGroup))
	gGroups.PUT(":id", inkstone.HandlerAdapter(updateGroup))
}

type groupReq struct {
	Type  int `json:"type" form:"type" binding:"required,eq=1|eq=2" example:"1"`
	AppId int `json:"appId" form:"appId" binding:"required,min=100000" example:"100000"`
}

type pagingGroupReq struct {
	inkstone.PagingRequest
	groupReq
}

type groupRes struct {
	inkstone.Response
	Name    string `json:"name,omitempty"`
	Type    int    `json:"type,omitempty"`
	AppId   int    `json:"appId,omitempty"`
	AppName string `json:"appName,omitempty"`
	Active  bool   `json:"active"`
}

// groups godoc
//
//	@Summary		Show groups
//	@Description	Show groups
//	@Tags			admin_group
//	@Router			/admin/groups	[get]
//	@Security		ApiKeyAuth
//	@Param			type	query		int	true	"type"
//	@Param			appId	query		int	true	"appId"
//	@Param			offset	query		int	false	"offset"
//	@Param			limit	query		int	true	"limit"
//	@Success		200		{object}	inkstone.PagingResponse[groupRes]
//	@Failure		400		{object}	inkstone.ClientError
//	@Failure		401		{object}	inkstone.ClientError
//	@Failure		403		{object}	inkstone.ClientError
//	@Failure		500		{string}	empty
func groups(c *inkstone.Context) {
	appCtx := c.AppContext()

	req := new(pagingGroupReq)
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var total int
	var groups []model.GroupWithApp

	if err := inkstone.Transaction(appCtx, func(tx *sqlx.Tx) (err error) {
		if total, err = orm.Group(appCtx).CountWithTx(req.Type, req.AppId, tx); err != nil {
			return
		}

		groups, err = orm.Group(appCtx).PaginationWithTx(req.Type, req.AppId, req.Offset, req.Limit, tx)
		return
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res []groupRes
	for i := range groups {
		group := &groups[i]
		res = append(res, groupRes{
			Response: inkstone.Response{
				Id:        int(group.Id),
				CreatedAt: group.CreatedAt,
				UpdatedAt: group.UpdatedAt,
			},
			Name:    group.Name,
			Type:    int(group.Type),
			AppId:   int(group.AppId),
			AppName: group.AppName,
			Active:  group.Active,
		})
	}

	c.Response(&inkstone.PagingResponse[groupRes]{
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  total,
		Items:  res,
	})
}

type addGroupReq struct {
	Name string `json:"name" binding:"required,min=2" example:"developer"`
	groupReq
}

// addGroup godoc
//
//	@Summary		Add a group
//	@Description	Add a group
//	@Tags			admin_group
//	@Router			/admin/groups	[post]
//	@Security		ApiKeyAuth
//	@Param			addGroupReq	body		addGroupReq	true	"request body"
//	@Success		200			{object}	groupRes
//	@Failure		400			{object}	inkstone.ClientError
//	@Failure		401			{object}	inkstone.ClientError
//	@Failure		403			{object}	inkstone.ClientError
//	@Failure		500			{string}	empty
func addGroup(c *inkstone.Context) {
	req := new(addGroupReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	group := model.NewGroup(req.Name, model.GroupType(req.Type), uint32(req.AppId))
	if err := orm.Group(c.AppContext()).Insert(group); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&groupRes{
		Response: inkstone.Response{
			Id: int(group.Id),
		},
		Name:   group.Name,
		Type:   int(group.Type),
		AppId:  int(group.AppId),
		Active: group.Active,
	})
}

type updateGroupParam struct {
	Id int `uri:"id" binding:"required,min=100000"`
}

type updateGroupReq struct {
	Name         string `json:"name" binding:"omitempty,min=2" example:"ceo"`
	ActiveToggle bool   `json:"activeToggle" example:"false"`
}

// updateGroup godoc
//
//	@Summary		Update a group
//	@Description	Update a group
//	@Tags			admin_group
//	@Router			/admin/groups/{id}	[put]
//	@Security		ApiKeyAuth
//	@Param			id				path		int				true	"group id"
//	@Param			updateGroupReq	body		updateGroupReq	true	"request body"
//	@Success		200				{object}	groupRes
//	@Failure		400				{object}	inkstone.ClientError
//	@Failure		401				{object}	inkstone.ClientError
//	@Failure		403				{object}	inkstone.ClientError
//	@Failure		500				{string}	empty
func updateGroup(c *inkstone.Context) {
	param := new(updateGroupParam)

	if err := c.ShouldBindUri(param); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	req := new(updateGroupReq)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(inkstone.ValidationNotAllFieldsZero, req)
	}

	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errors.ERR_BAD_REQUEST)
		return
	}

	var (
		appCtx = c.AppContext()
		group  *model.Group
	)

	if err := inkstone.Transaction(appCtx, func(tx *sqlx.Tx) (err error) {
		group, err = orm.Group(appCtx).GetWithTx(param.Id, tx)
		if err != nil {
			return
		}

		if req.Name == group.Name {
			return errs.New("group's name not changed")
		} else if req.Name != "" {
			group.Name = req.Name
		}
		if req.ActiveToggle {
			group.Active = !group.Active
		}

		return orm.Group(appCtx).UpdateWithTx(group, tx)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&groupRes{
		Response: inkstone.Response{
			Id: int(group.Id),
		},
		Name:   group.Name,
		Type:   int(group.Type),
		AppId:  int(group.AppId),
		Active: group.Active,
	})
}
