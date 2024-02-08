package ext

import "fmt"

type ClientError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ClientError) Error() string {
	return fmt.Sprintf("%s(%s)", e.Message, e.Code)
}

var _ error = (*ClientError)(nil)

var (
	ERR_BAD_REQUEST error = &ClientError{
		Code: "BadRequest",
	}

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

	ERR_MISSING_ACCESS_TOKEN error = &ClientError{
		Code: "MissingAccessToken",
	}

	ERR_REVOKED_ACCESS_TOKEN error = &ClientError{
		Code: "RevokedAccessToken",
	}

	ERR_EXPIRED_ACCESS_TOKEN error = &ClientError{
		Code: "ExpiredAccessToken",
	}
)
