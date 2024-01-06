package store

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	IsEmailExists(email string) (bool, error)
}

type MongoUserStore struct {
	collection *mongo.Collection
}

func NewMongoUserStore(db *mongo.Database) *MongoUserStore {
	return &MongoUserStore{
		collection: db.Collection("user"),
	}
}

func (m *MongoUserStore) IsEmailExists(email string) (bool, error) {
	filter := bson.D{{"email", email}}
	err := m.collection.FindOne(nil, filter).Err()

	if err == nil {
		return true, nil
	} else if err == mongo.ErrNoDocuments {
		return false, nil
	}

	return false, err
}