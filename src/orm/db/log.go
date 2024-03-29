// Package db Code generated by authink/orm. DO NOT EDIT
package db

import (
	"github.com/authink/ink/src/orm/models"
	"github.com/authink/orm/db"
	sbd "github.com/authink/sqlbuilder"
)

type log struct {
	Detail sbd.Field
}

// Tname implements db.Table.
func (*log) Tname() sbd.Table {
	return "s_logs"
}

var _ db.Table = (*log)(nil)

var Log log

func init() {
	db.Bind(&Log, &models.Log{})
}
