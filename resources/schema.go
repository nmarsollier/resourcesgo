package resources

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resource struct {
	ID       primitive.ObjectID `bson:"_id"`
	Project  string             `bson:"project" validate:"required,min=1,max=20"`
	Language string             `bson:"language" validate:"required,min=1,max=20"`
	SemVer   string             `bson:"semVer" validate:"required,min=5,max=50"`
	Values   map[string]string  `bson:"values"`
	Created  time.Time          `bson:"created"`
	Enabled  bool               `bson:"enabled"`
}

func newResource(
	project string,
	language string,
	semVer string,
	values map[string]string,
) *Resource {
	return &Resource{
		ID:       primitive.NewObjectID(),
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
