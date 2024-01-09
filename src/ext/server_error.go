package ext

import "errors"

type ServerError struct {
	err error
}

func (e *ServerError) Error() string {
	return e.err.Error()
}

var _ error = (*ServerError)(nil)

var (
	ERR_SRV_DB_TIMEOUT error = &ServerError{
		errors.New("MySQL connect timeout"),
	}
)
