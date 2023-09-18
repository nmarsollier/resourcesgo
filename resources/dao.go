package resources

import (
	"context"

	"github.com/nmarsollier/resourcesgo/tools/db"
	"github.com/nmarsollier/resourcesgo/tools/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	collection = database.Collection("resources")

	return collection, nil
}

func insert(
	project string,
	language string,
	semVer string,
	values map[string]string,
) (*Resource, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	resource := newResource(project, language, semVer, values)

	if err = resource.ValidateSchema(); err != nil {
		return nil, err
	}

	if _, err = collection.InsertOne(context.Background(), resource); err != nil {
		return nil, err
	}

	return resource, nil
}

func findByID(id string) (*Resource, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.ErrID
	}

	resource := &Resource{}
	filter := bson.M{"_id": _id, "enabled": true}

	if err = collection.FindOne(context.Background(), filter).Decode(resource); err != nil {
		return nil, err
	}

	return resource, nil
}

func findBy(project string, language string, semVer string) (*Resource, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	resource := &Resource{}
	filter := bson.M{"project": project, "language": language, "semVer": semVer, "enabled": true}

	if err = collection.FindOne(context.Background(), filter).Decode(resource); err != nil {
		return nil, err
	}

	return resource, nil
}

type semVerResult struct {
	semVer string
}

func findVersions(project string, language string) ([]string, error) {
	collection, err := dbCollection()
	if err != nil {
		return nil, err
	}

	projection := bson.D{
		{Key: "semVer", Value: 1},
	}

	filter := bson.M{"project": project, "language": language, "enabled": true}

	opts := options.Find().SetProjection(projection)

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	var result []string

	for cursor.Next(context.Background()) {
		var value string = cursor.Current.Lookup("semVer").StringValue()
		result = append(result, value)
	}

	return result, nil
}

func delete(id primitive.ObjectID) error {
	collection, err := dbCollection()
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": bson.M{
			"enabled": false,
		}},
	)

	return err
}
