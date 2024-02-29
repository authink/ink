package orm

import (
	"database/sql"
	"errors"

	"github.com/authink/inkstone"
)

func handleSaveResult(result sql.Result, model *inkstone.Model) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	} else if lastId > 0 {
		model.Id = uint32(lastId)
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

func handleInsertResult(result sql.Result, model *inkstone.Model) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	model.Id = uint32(lastId)
	return
}
