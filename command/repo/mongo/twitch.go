package command

import (
	"context"
	"phillzbot/domain"

	"go.mongodb.org/mongo-driver/bson"
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

	if nil == err {
		return results, err
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

func (m *MongoCommandRepo) FetchById(ctx context.Context, id string) (data domain.TwitchCommand, err error) {
	return data, nil
}
func (m *MongoCommandRepo) FetchByName(ctx context.Context, name string) (data domain.TwitchCommand, err error) {
	return data, nil

}
func (m *MongoCommandRepo) Update(ctx context.Context, tc *domain.TwitchCommand) error {
	return nil
}
func (m *MongoCommandRepo) Create(ctx context.Context, tc *domain.TwitchCommand) error {
	return nil
}
func (m *MongoCommandRepo) Delete(ctx context.Context, id string) error {
	return nil

}
