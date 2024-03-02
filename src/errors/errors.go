package errors

import "github.com/authink/inkstone/web"

var (
	ERR_BAD_REQUEST error = &web.ClientError{
		Code: "BadRequest",
	}

	ERR_INVALID_APP error = &web.ClientError{
		Code: "InvalidApp",
	}

	ERR_UNSUPPORTED_APP error = &web.ClientError{
		Code: "UnsupportedApp",
	}

	ERR_INVALID_ACCOUNT error = &web.ClientError{
		Code: "InvalidEmailOrPassword",
	}

	ERR_INVALID_REFRESH_TOKEN error = &web.ClientError{
		Code: "InvalidRefreshToken",
	}

	ERR_INVALID_ACCESS_TOKEN error = &web.ClientError{
		Code: "InvalidAccessToken",
	}

	ERR_MISSING_ACCESS_TOKEN error = &web.ClientError{
		Code: "MissingAccessToken",
	}

	ERR_REVOKED_ACCESS_TOKEN error = &web.ClientError{
		Code: "RevokedAccessToken",
	}

	ERR_EXPIRED_ACCESS_TOKEN error = &web.ClientError{
		Code: "ExpiredAccessToken",
	}

	ERR_NO_PERMISSION error = &web.ClientError{
		Code: "NoPermission",
	}
)
