package orm

import (
	"database/sql"
	"errors"

	"github.com/authink/inkstone/model"
)

type resultHandlerFunc func(sql.Result, model.Identifier) error

type dbExecutor interface {
	NamedExec(string, any) (sql.Result, error)
}

func namedExec(executor dbExecutor, statement string, m model.Identifier, handleResult resultHandlerFunc) (err error) {
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

func handleSaveResult(result sql.Result, m model.Identifier) (err error) {
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

func handleInsertResult(result sql.Result, m model.Identifier) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	m.SetId(uint32(lastId))
	return
}
