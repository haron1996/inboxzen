package apperror

import "fmt"

type APPError struct {
	Message string
	Code    int
	Err     error
}

func (e *APPError) Error() string {
	return fmt.Sprintf("message: %s code: %d error: %v", e.Message, e.Code, e.Err)
}
