package projects

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Project struct {
	ID      string    `db:"id" json:"id"`
	Name    string    `db:"name" json:"name" validate:"required,min=1,max=20"`
	Created time.Time `db:"created" json:"created"`
	Enabled bool      `db:"enabled" json:"enabled"`
}

func newProject(id string, name string) *Project {
	return &Project{
		ID:      id,
		Enabled: true,
		Created: time.Now(),
		Name:    name,
	}
}

func (e *Project) ValidateSchema() error {
	return validator.New().Struct(e)
}
