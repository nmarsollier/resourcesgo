package resources

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Resource struct {
	ID         string            `db:"id"`
	ProjectID  string            `db:"project" validate:"required,min=1,max=20"`
	LanguageID string            `db:"language" validate:"required,min=1,max=20"`
	SemVer     string            `db:"sem_ver" validate:"required,min=5,max=50"`
	Values     map[string]string `db:"values"`
	Created    time.Time         `db:"created"`
	Enabled    bool              `db:"enabled"`
}

func NewResource(
	projectID string,
	languageID string,
	semVer string,
	values map[string]string,
) *Resource {
	return &Resource{
		ID:         uuid.New().String(),
		ProjectID:  projectID,
		LanguageID: languageID,
		SemVer:     semVer,
		Values:     values,
		Enabled:    true,
		Created:    time.Now(),
	}
}

func (e *Resource) ValidateSchema() error {
	return validator.New().Struct(e)
}
