package domain

import (
	"context"
)

type DiscordRole int

const (
	GodTier DiscordRole = iota
	Moderators
	Homies
	TwitchSubs
	FightClub
	Streamers
	Creatives
	Developers
	Gamers
	TheBlock
)

func (role DiscordRole) String() string {
	return [...]string{
		"GOD TIER",
		"MODERATORS",
		"THE HOMIES",
		"SUBSCRIBERS",
		"FIGHT CLUB",
		"STREAMERS",
		"CREATIVES",
		"DEVELOPERS",
		"GAMERS",
		"THE BLOCK"}[role]
}

type DiscordCommand struct {
	Id          string      `json:"id" bson:"id"`
	Name        string      `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	Role        DiscordRole `json:"role" bson:"role"`
	Value       string      `json:"value" bson:"value"`
	Cooldown    int         `json:"cooldown" bson:"cooldown"`
}

type DiscordService interface {
	FetchCommand(ctx context.Context) (data []DiscordCommand, err error)
	FetchCommandById(ctx context.Context, id string) (DiscordCommand, error)
	FetchCommandByName(ctx context.Context, name string) (DiscordCommand, error)
	UpdateCommand(ctx context.Context, tc DiscordCommand) error
	CreateCommand(ctx context.Context, tc DiscordCommand) error
	DeleteCommand(ctx context.Context, id string) error
}

type DiscordRepo interface {
	FetchCommand(ctx context.Context) (data []DiscordCommand, err error)
	FetchCommandById(ctx context.Context, id string) (DiscordCommand, error)
	FetchCommandByName(ctx context.Context, name string) (DiscordCommand, error)
	UpdateCommand(ctx context.Context, dc DiscordCommand) error
	CreateCommand(ctx context.Context, dc DiscordCommand) error
	DeleteCommand(ctx context.Context, id string) error
}
