package orm

import (
	libsql "database/sql"
	"errors"

	"github.com/authink/inkstone/orm"
)

type resultHandlerFunc func(libsql.Result, orm.Identifier) error

type dbExecutor interface {
	NamedExec(string, any) (libsql.Result, error)
}

func namedExec(executor dbExecutor, statement string, m orm.Identifier, handleResult resultHandlerFunc) (err error) {
	result, err := executor.NamedExec(
		statement,
		m,
	)
	if err != nil {
		return
	}

	if handleResult != nil {
		err = handleResult(result, m)
	}
	return
}

func handleSaveResult(result libsql.Result, m orm.Identifier) (err error) {
	if err = handleInsertResult(result, m); err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	} else if rowsAffected == 0 {
		err = errors.New("duplicate key")
	}
	return
}

func handleInsertResult(result libsql.Result, m orm.Identifier) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	m.SetId(uint32(lastId))
	return
}
