package errs

import (
	"github.com/nmarsollier/resourcesgo/tools/strs"
)

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
