package errs

import "github.com/nmarsollier/resourcesgo/tools/strs"

type Validation struct {
	Messages []Field `json:"messages"`
}

type Field struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

func NewValidationField(field string, err string) *Validation {
	return &Validation{
		Messages: []Field{
			{
				Path:    field,
				Message: err,
			},
		},
	}
}

func NewValidation() *Validation {
	return &Validation{
		Messages: []Field{},
	}
}

func (e *Validation) Error() string {
	return strs.ToJson(e)
}

func (e *Validation) Add(path string, message string) *Validation {
	err := Field{
		Path:    path,
		Message: message,
	}
	e.Messages = append(e.Messages, err)
	return e
}

func (e *Validation) Size() int {
	return len(e.Messages)
}
