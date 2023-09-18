package projects

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name" validate:"required,min=1,max=20"`
	Created time.Time          `bson:"created"`
	Enabled bool               `bson:"enabled"`
}

func newProject(name string) *Project {
	return &Project{
		ID:      primitive.NewObjectID(),
		Enabled: true,
		Created: time.Now(),
		Name:    name,
	}
}

func (e *Project) ValidateSchema() error {
	return validator.New().Struct(e)
}
