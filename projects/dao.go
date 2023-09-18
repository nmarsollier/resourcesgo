package projects

import (
	"context"

	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func dbCollection() (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	database, err := db.Get()
	if err != nil {
		return nil, err
	}

	collection = database.Collection("projects")

	return collection, nil
}

func insert(name string) (*Project, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	project := newProject(name)

	if err = project.ValidateSchema(); err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(context.Background(), project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func findByID(id string) (*Project, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.ErrID
	}

	project := &Project{}
	filter := bson.M{"_id": _id}

	if err = collection.FindOne(context.Background(), filter).Decode(project); err != nil {
		return nil, err
	}

	return project, nil
}

func delete(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.ErrID
	}

	collection, err := dbCollection()
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": _id},
		bson.M{"$set": bson.M{
			"enabled": false,
		}},
	)

	return err
}
