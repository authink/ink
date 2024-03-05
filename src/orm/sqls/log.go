package sqls

import (
	"fmt"

	"github.com/authink/ink.go/src/orm/db"
	"github.com/authink/inkstone/orm/sql"
)

type log interface {
	sql.Inserter
	sql.Finder
}

type logImpl struct{}

// Find implements log.
func (l *logImpl) Find() string {
	return fmt.Sprintf("SELECT id, created_at, detail FROM %s ORDER BY id DESC", db.Log.Tname())
}

// Insert implements log.
func (l *logImpl) Insert() string {
	return fmt.Sprintf("INSERT INTO %s (detail) VALUES (:detail)", db.Log.Tname())
}

var Log log = new(logImpl)
