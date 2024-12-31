package resources

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Resource struct {
	ID       string            `db:"id" json:"id"`
	Project  string            `db:"project" json:"project" validate:"required,min=1,max=20"`
	Language string            `db:"language" json:"language" validate:"required,min=1,max=20"`
	SemVer   string            `db:"sem_ver" json:"semVer" validate:"required,min=5,max=50"`
	Values   map[string]string `db:"values" json:"values"`
	Created  time.Time         `db:"created" json:"created"`
	Enabled  bool              `db:"enabled" json:"enabled"`
}

func newResource(
	project string,
	language string,
	semVer string,
	values map[string]string,
) *Resource {
	return &Resource{
		ID:       uuid.New().String(),
		Project:  project,
		Language: language,
		SemVer:   semVer,
		Values:   values,
		Enabled:  true,
		Created:  time.Now(),
	}
}

func (e *Resource) ValidateSchema() error {
	return validator.New().Struct(e)
}
