package models

import "github.com/authink/inkstone/orm/model"

type GroupType uint32

const (
	GROUP_TYPE_ROLE GroupType = iota + 1
	GROUP_TYPE_RESOURCE
)

type Group struct {
	model.Base
	Name   string
	Type   uint32
	AppId  uint32 `db:"app_id"`
	Active bool
}

func NewGroup(name string, gtype GroupType, appId uint32) *Group {
	return &Group{
		Name:   name,
		Type:   uint32(gtype),
		AppId:  appId,
		Active: true,
	}
}

func (g *Group) IsTypeRole() bool {
	return g.Type == uint32(GROUP_TYPE_ROLE)
}

func (g *Group) IsTypeResource() bool {
	return g.Type == uint32(GROUP_TYPE_RESOURCE)
}

type GroupWithApp struct {
	Group
	AppName string `db:"app_name"`
}

type GroupPage struct {
	model.Page
	Type  int
	AppId int `db:"app_id"`
}
