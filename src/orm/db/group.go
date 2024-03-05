package db

import (
	"github.com/authink/ink.go/src/orm/models"
	"github.com/authink/inkstone/orm/db"
)

type group struct {
	Name   string
	Type   string
	AppId  string
	Active string
}

// Tname implements db.Table.
func (a *group) Tname() string {
	return "s_groups"
}

var _ db.Table = (*group)(nil)

var Group group

type groupWithApp struct {
	group
	AppName string
}

var GroupWithApp groupWithApp

func init() {
	db.Bind(&Group, &models.Group{})
	db.Bind(&GroupWithApp, &models.GroupWithApp{})
}
