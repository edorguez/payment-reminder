package errors

import "fmt"

var (
	ErrGeneral      = fmt.Errorf("GENERAL")
	ErrNotFound     = fmt.Errorf("NOT_FOUND")
	ErrConflict     = fmt.Errorf("CONFLIC")
	ErrInvalidInput = fmt.Errorf("INVALID_INPUT")
)

type Error struct {
	Err     error
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

// func (e *Error) Is(target error) bool {
// 	t, ok := target.(*Error)
// 	if !ok {
// 		return false
// 	}
// 	return e.Err == t.Err
// }
