package repository

import (
	"context"
	"log"
	"phillzbot/domain"
)

const (
	COLLECTION = "commands"
)

type CommandRepository interface {
	SaveCommand(command *domain.TwitchCommand) (*domain.TwitchCommand, error)
	GetAll() ([]domain.TwitchCommand, error)
	GetById(commandId string) (*domain.TwitchCommand, error)
	GetByName(name string) (*domain.TwitchCommand, error)
	Delete(command *domain.TwitchCommand) error
}

type commandRepository struct {
	m *MongoDB
}

func (cr *commandRepository) SaveCommand(command *domain.TwitchCommand) (*domain.TwitchCommand, error) {
	// mongodb, err := GetConnection()

	collection := cr.m.Database.Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), command)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return command, nil
}

func GetAll() ([]domain.TwitchCommand, error) {
	return nil, nil
}
func GetById(commandId string) (*domain.TwitchCommand, error) {
	return nil, nil
}
func GetByName(name string) (*domain.TwitchCommand, error) {
	return nil, nil
}
func Delete(command *domain.TwitchCommand) error {
	return nil
}
