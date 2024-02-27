package errors

import "github.com/authink/inkstone"

var (
	ERR_BAD_REQUEST error = &inkstone.ClientError{
		Code: "BadRequest",
	}

	ERR_INVALID_APP error = &inkstone.ClientError{
		Code: "InvalidApp",
	}

	ERR_UNSUPPORTED_APP error = &inkstone.ClientError{
		Code: "UnsupportedApp",
	}

	ERR_INVALID_ACCOUNT error = &inkstone.ClientError{
		Code: "InvalidEmailOrPassword",
	}

	ERR_INVALID_REFRESH_TOKEN error = &inkstone.ClientError{
		Code: "InvalidRefreshToken",
	}

	ERR_INVALID_ACCESS_TOKEN error = &inkstone.ClientError{
		Code: "InvalidAccessToken",
	}

	ERR_MISSING_ACCESS_TOKEN error = &inkstone.ClientError{
		Code: "MissingAccessToken",
	}

	ERR_REVOKED_ACCESS_TOKEN error = &inkstone.ClientError{
		Code: "RevokedAccessToken",
	}

	ERR_EXPIRED_ACCESS_TOKEN error = &inkstone.ClientError{
		Code: "ExpiredAccessToken",
	}

	ERR_NO_PERMISSION error = &inkstone.ClientError{
		Code: "NoPermission",
	}
)
