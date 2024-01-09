package err

type ClientError struct {
	Code    string
	Message string
}

func (e *ClientError) Error() string {
	return e.Message
}

var _ error = (*ClientError)(nil)

var (
	ERR_CLI_BAD_EMAIL error = &ClientError{
		Code: "invalid.email", Message: "Bad email format.",
	}
)
