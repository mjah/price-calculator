package errors

import (
	"fmt"
)

// ErrorCode stores the error code along with a  message
type ErrorCode struct {
	code    int
	message string
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.message)
}

// New creates and returns a new instance of ErrorCode
func New(code int, message string) *ErrorCode {
	return &ErrorCode{
		code:    code,
		message: message,
	}
}

// ErrorCode Constants
const (
	NegDenUncappedFees int = 0
	NegDenCappedFees   int = 1
)
