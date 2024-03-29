// Package db Code generated by authink/orm. DO NOT EDIT
package db

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/orm/db"
	sbd "github.com/authink/sqlbuilder"
)

type group struct {
	Name sbd.Field

	Type sbd.Field

	AppId sbd.Field

	Active sbd.Field
}

// Tname implements db.Table.
func (*group) Tname() sbd.Table {
	return "s_groups"
}

var _ db.Table = (*group)(nil)

var Group group

func init() {
	db.Bind(&Group, &models.Group{})
}
