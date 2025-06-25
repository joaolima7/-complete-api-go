package errs

type ErrorCode string

var (
	ErrDomainValidation ErrorCode = "DOMAIN_VALIDATION"
	ErrConflict         ErrorCode = "CONFLICT"
	ErrNotFound         ErrorCode = "NOT_FOUND"

	ErrInternal ErrorCode = "INTERNAL"
)

type AppError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return string(e.Code) + ": " + e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func DomainValidation(message string, err error) *AppError {
	return New(ErrDomainValidation, message, err)
}

func Conflict(message string, err error) *AppError {
	return New(ErrConflict, message, err)
}

func NotFound(message string, err error) *AppError {
	return New(ErrNotFound, message, err)
}

func Internal(message string, err error) *AppError {
	return New(ErrInternal, message, err)
}
