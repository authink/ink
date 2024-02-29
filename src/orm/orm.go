package orm

import (
	libsql "database/sql"
	"errors"

	"github.com/authink/inkstone"
)

type resultHandlerFunc func(libsql.Result, inkstone.ModelOf) error

type dbExecutor interface {
	NamedExec(query string, arg interface{}) (libsql.Result, error)
}

func namedExec(executor dbExecutor, statement string, m inkstone.ModelOf, handleResult resultHandlerFunc) (err error) {
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

func handleSaveResult(result libsql.Result, m inkstone.ModelOf) (err error) {
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

func handleInsertResult(result libsql.Result, m inkstone.ModelOf) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	m.Of().Id = uint32(lastId)
	return
}
