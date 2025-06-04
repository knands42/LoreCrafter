package utils

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Errors []string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation errors: %v", e.Errors)
}

func (e *ValidationError) Is(target error) bool {
	var validationError *ValidationError
	ok := errors.As(target, &validationError)
	return ok
}
