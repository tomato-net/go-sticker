package alphavantage

import (
	"fmt"
)

type Error interface {
	Error() error
}

// TODO: Add additional error handling for Information key when not-premium endpoint
type ErrorMessage struct {
	Message string `json:"Error Message,omitempty"`
}

func (e ErrorMessage) Error() error {
	if e.Message == "" {
		return nil
	}

	return fmt.Errorf(e.Message)
}

func IsError(e Error) bool {
	return e.Error() != nil
}
