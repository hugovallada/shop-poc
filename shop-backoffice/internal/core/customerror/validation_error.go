package customerror

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	validationErrorsMessage map[string][]string
}

func (ve ValidationError) Error() string {
	var baseMessage string = "Validation has failed for field "
	if len(ve.validationErrorsMessage) > 1 {
		baseMessage = "Validation has failed for fields"
	}
	var keys []string
	for key := range ve.validationErrorsMessage {
		keys = append(keys, key)
	}
	return fmt.Sprintf("%s: %s\n %v", baseMessage, strings.Join(keys, ","), ve.validationErrorsMessage)
}
