package command

import (
	"context"
	"fmt"
	"phillzbot/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCommandRepo struct {
	db         *mongo.Database
	collection string
}

func NewMongoCommandRepo(db *mongo.Database) domain.TwitchCommandRepo {
	return &MongoCommandRepo{db: db, collection: "twitch_commands"}
}

func (m *MongoCommandRepo) Fetch(ctx context.Context) (data []domain.TwitchCommand, err error) {
	val, err := m.db.Collection(m.collection).Find(ctx, bson.D{{}})
	var results = []domain.TwitchCommand{}

	if err != nil {
		return nil, err
	}

	for val.Next(ctx) {
		c := domain.TwitchCommand{}
		err := val.Decode(&c)

		if err != nil {
			return results, err
		}
		results = append(results, c)
	}

	if len(results) == 0 {
		return results, mongo.ErrNoDocuments
	}
	return results, nil
}

func (m *MongoCommandRepo) FetchById(ctx context.Context, id string) (domain.TwitchCommand, error) {
	var command domain.TwitchCommand

	reqId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": reqId}
	err := m.db.Collection(m.collection).FindOne(ctx, filter).Decode(&command)

	if err != nil {
		return command, err
	}

	return command, nil
}
func (m *MongoCommandRepo) FetchByName(ctx context.Context, name string) (data domain.TwitchCommand, err error) {
	return data, nil

}
func (m *MongoCommandRepo) Update(ctx context.Context, tc *domain.TwitchCommand, cmdId primitive.ObjectID) error {

	filter := bson.M{"_id": cmdId}
	_, err := m.db.Collection(m.collection).UpdateOne(context.Background(), filter, bson.M{"$set": tc})

	if err != nil {
		return err
	}

	return nil
}
func (m *MongoCommandRepo) Create(ctx context.Context, tc *domain.TwitchCommand) error {
	result, err := m.db.Collection(m.collection).InsertOne(ctx, tc)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}
func (m *MongoCommandRepo) Delete(ctx context.Context, id string) error {
	return nil
}
