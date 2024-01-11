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
	ERR_INVALID_APP error = &ClientError{
		Code: "InvalidApp",
	}

	ERR_UNSUPPORTED_APP error = &ClientError{
		Code: "UnsupportedApp",
	}

	ERR_INVALID_ACCOUNT error = &ClientError{
		Code: "InvalidEmailOrPassword",
	}

	ERR_INVALID_REFRESH_TOKEN error = &ClientError{
		Code: "InvalidRefreshToken",
	}

	ERR_INVALID_ACCESS_TOKEN error = &ClientError{
		Code: "InvalidAccessToken",
	}
)
