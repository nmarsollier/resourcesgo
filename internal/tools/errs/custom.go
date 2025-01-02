package errs

import (
	"github.com/nmarsollier/resourcesgo/internal/tools/strs"
)

// Custom represents a custom error with an HTTP status code and a message.
// The status field holds the HTTP status code, and the Message field holds
// the error message to be returned in JSON format.
type Custom struct {
	status  int
	Message string `json:"error"`
}

func NewCustom(status int, message string) *Custom {
	return &Custom{
		status:  status,
		Message: message,
	}
}

func (e *Custom) Error() string {
	return strs.ToJson(e.Message)
}

func (e *Custom) Status() int {
	return e.status
}
