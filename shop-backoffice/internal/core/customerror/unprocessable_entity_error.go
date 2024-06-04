package customerror

import "fmt"

type UnprocessableEntityError struct {
	message string
}

func NewUnprocessableEntityError(message string) UnprocessableEntityError {
	return UnprocessableEntityError{
		message: message,
	}
}

func (ue UnprocessableEntityError) Error() string {
	return fmt.Sprintf("Can't process over entity: %s", ue.message)
}
