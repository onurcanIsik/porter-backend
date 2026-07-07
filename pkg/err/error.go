package apperr

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) StatusCode() int {
	return e.Code
}

func NewBadRequestError(message string) *Error {
	return NewError(400, message)
}

func NewUnauthorizedError(message string) *Error {
	return NewError(401, message)
}

func NewForbiddenError(message string) *Error {
	return NewError(403, message)
}

func NewNotFoundError(message string) *Error {
	return NewError(404, message)
}

func NewInternalServerError(message string) *Error {
	return NewError(500, message)
}

func NewConflictError(message string) *Error {
	return NewError(409, message)
}

func NewTooManyRequestsError(message string) *Error {
	return NewError(429, message)
}
