package errors

import "fmt"

var (
	ErrGeneral      = fmt.Errorf("GENERAL")
	ErrNotFound     = fmt.Errorf("NOT_FOUND")
	ErrConflict     = fmt.Errorf("CONFLIC")
	ErrInvalidInput = fmt.Errorf("INVALID_INPUT")
	ErrPublishEvent = fmt.Errorf("PUBLISH_EVENT")
	ErrConsumeEvent = fmt.Errorf("CONSUME_EVENT")
	ErrFirebase     = fmt.Errorf("ERROR_FIREBASE")
)

type Error struct {
	Err     error
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}
