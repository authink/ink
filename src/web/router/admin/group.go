package admin

import (
	"errors"

	"github.com/authink/ink/src/authz"
	"github.com/authink/ink/src/orm"
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/ink/src/web/errs"
	"github.com/authink/ink/src/web/middleware"
	"github.com/authink/orm/model"
	"github.com/authink/stone/web"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func setupGroupGroup(gAdmin *gin.RouterGroup) {
	gGroups := gAdmin.Group(authz.Groups.Name)
	gGroups.Use(
		middleware.Authz(authz.Groups),
		middleware.Log(authz.Groups),
	)
	gGroups.GET("", web.HandlerAdapter(groups))
	gGroups.POST("", web.HandlerAdapter(addGroup))
	gGroups.PUT(":id", web.HandlerAdapter(updateGroup))
}

type groupReq struct {
	Type  int `json:"type" form:"type" binding:"required,eq=1|eq=2" example:"1"`
	AppId int `json:"appId" form:"appId" binding:"required,min=100000" example:"100000"`
}

type pagingGroupReq struct {
	web.PagingRequest
	groupReq
}

type groupRes struct {
	web.Response
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
//	@Success		200		{object}	web.PagingResponse[groupRes]
//	@Failure		400		{object}	web.ClientError
//	@Failure		401		{object}	web.ClientError
//	@Failure		403		{object}	web.ClientError
//	@Failure		500		{string}	empty
func groups(c *web.Context) {
	appCtx := c.AppContext()

	req := &pagingGroupReq{}
	if err := c.ShouldBindQuery(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	var total int
	var groups []models.GroupWithApp

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		groupPage := models.GroupPage{
			Page: model.Page{
				Offset: req.Offset,
				Limit:  req.Limit,
			},
			Type:  req.Type,
			AppId: req.AppId,
		}

		if total, err = orm.Group(appCtx).CountTx(tx, &groupPage); err != nil {
			return
		}

		groups, err = orm.Group(appCtx).PaginationTx(tx, &groupPage)
		return
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	var res = []groupRes{}
	for i := range groups {
		group := &groups[i]
		res = append(res, groupRes{
			Response: web.Response{
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

	c.Response(&web.PagingResponse[groupRes]{
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
//	@Failure		400			{object}	web.ClientError
//	@Failure		401			{object}	web.ClientError
//	@Failure		403			{object}	web.ClientError
//	@Failure		500			{string}	empty
func addGroup(c *web.Context) {
	req := &addGroupReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	group := models.NewGroup(req.Name, uint16(req.Type), uint32(req.AppId))
	if err := orm.Group(c.AppContext()).Insert(group); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&groupRes{
		Response: web.Response{
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
//	@Failure		400				{object}	web.ClientError
//	@Failure		401				{object}	web.ClientError
//	@Failure		403				{object}	web.ClientError
//	@Failure		500				{string}	empty
func updateGroup(c *web.Context) {
	param := &updateGroupParam{}

	if err := c.ShouldBindUri(param); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	req := &updateGroupReq{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.AbortWithClientError(errs.ERR_BAD_REQUEST)
		return
	}

	var (
		appCtx = c.AppContext()
		group  models.Group
	)
	group.Id = uint32(param.Id)

	if err := appCtx.Transaction(func(tx *sqlx.Tx) (err error) {
		err = orm.Group(appCtx).GetTx(tx, &group)
		if err != nil {
			return
		}

		if req.Name == group.Name {
			return errors.New("group's name not changed")
		} else if req.Name != "" {
			group.Name = req.Name
		}
		if req.ActiveToggle {
			group.Active = !group.Active
		}

		return orm.Group(appCtx).UpdateTx(tx, &group)
	}); err != nil {
		c.AbortWithServerError(err)
		return
	}

	c.Response(&groupRes{
		Response: web.Response{
			Id: int(group.Id),
		},
		Name:   group.Name,
		Type:   int(group.Type),
		AppId:  int(group.AppId),
		Active: group.Active,
	})
}
