package orm

import (
	"database/sql"
	"errors"

	"github.com/authink/inkstone/model"
)

type afterExecFunc func(sql.Result, model.Identifier) error

type dbExecutor interface {
	Exec(string, ...any) (sql.Result, error)
	NamedExec(string, any) (sql.Result, error)
	Get(any, string, ...any) error
	Select(any, string, ...any) error
}

func namedExec(executor dbExecutor, statement string, m model.Identifier, afterExec afterExecFunc) (err error) {
	result, err := executor.NamedExec(
		statement,
		m,
	)
	if err != nil {
		return
	}

	if afterExec != nil {
		err = afterExec(result, m)
	}
	return
}

func afterSave(result sql.Result, m model.Identifier) (err error) {
	if err = afterInsert(result, m); err != nil {
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

func afterInsert(result sql.Result, m model.Identifier) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	m.SetId(uint32(lastId))
	return
}

func get(executor dbExecutor, dest any, statement string, args ...any) error {
	return executor.Get(
		dest,
		statement,
		args...,
	)
}

func doSelect(executor dbExecutor, list any, statement string, args ...any) error {
	return executor.Select(
		list,
		statement,
		args...,
	)
}

func delete(executor dbExecutor, statement string, args ...any) (err error) {
	result, err := executor.Exec(
		statement,
		args...,
	)
	if err != nil {
		return
	}
	afterDelete(result)
	return
}

func afterDelete(result sql.Result) error {
	return nil
}
