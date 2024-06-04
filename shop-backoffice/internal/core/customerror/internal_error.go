package customerror

import "fmt"

type InternalError struct {
	message string
}

func (ie InternalError) Error() string {
	return fmt.Sprintf("Internal Application Error: %s. Contate o administrador do sistema ou suporte!", ie.message)
}

func NewInternalError(message string) InternalError {
	return InternalError{
		message: message,
	}
}
