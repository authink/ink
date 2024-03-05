package models

import "github.com/authink/inkstone/orm/model"

// @model
// @db s_groups
type Group struct {
	model.Base
	Name   string
	Type   uint16
	AppId  uint32 `db:"app_id"`
	Active bool
}

func NewGroup(name string, gtype uint16, appId uint32) *Group {
	return &Group{
		Name:   name,
		Type:   gtype,
		AppId:  appId,
		Active: true,
	}
}

// @model
// @embed Group
type GroupWithApp struct {
	Group
	AppName string `db:"app_name"`
}

// @param
type GroupPage struct {
	model.Page
	Type  int
	AppId int `db:"app_id"`
}
