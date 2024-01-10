package ext

import "fmt"

type ClientError struct {
	Code    string
	Message string
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("%s(%s)", e.Message, e.Code)
}

var _ error = (*ClientError)(nil)

var (
	ERR_CLI_INVALID_APP error = &ClientError{
		Code: "InvalidApp",
	}

	ERR_CLI_UNSUPPORTED_APP error = &ClientError{
		Code: "UnsupportedApp",
	}

	ERR_CLI_INVALID_ACCOUNT error = &ClientError{
		Code: "InvalidEmailOrPassword",
	}
)
