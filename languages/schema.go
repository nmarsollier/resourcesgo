package languages

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Language struct {
	ID      string    `db:"id" json:"id"  validate:"required,min=2,max=6"`
	Name    string    `db:"name" json:"name" validate:"required,min=1,max=20"`
	Created time.Time `db:"created" json:"created"`
	Enabled bool      `db:"enabled" json:"enabled"`
}

func newLanguage(id string, name string) *Language {
	return &Language{
		ID:      id,
		Enabled: true,
		Created: time.Now(),
		Name:    name,
	}
}

func (e *Language) ValidateSchema() error {
	return validator.New().Struct(e)
}
