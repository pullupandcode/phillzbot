package domain

import (
	"context"

	"github.com/Adeithe/go-twitch/api/helix"
	"github.com/Adeithe/go-twitch/api/kraken"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/Adeithe/go-twitch/pubsub"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TwitchCommandRole int

const (
	Broadcaster TwitchCommandRole = iota
	Moderator
	VIP
	Subscribers
	Follower
	Viewer
)

func (role TwitchCommandRole) String() string {
	return [...]string{"BROADCASTER", "MODERATOR", "VIP", "SUBSCRIBERS", "FOLLOWERS", "VIEWERS"}[role]
}

type TwitchCommand struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Role        TwitchCommandRole  `json:"role" bson:"role"`
	Value       string             `json:"value" bson:"value"`
	Cooldown    int                `json:"cooldown" bson:"cooldown"`
}

type TwitchCommandService interface {
	Fetch(ctx context.Context) (data []TwitchCommand, err error)
	FetchById(ctx context.Context, id string) (TwitchCommand, error)
	FetchByName(ctx context.Context, name string) (TwitchCommand, error)
	Update(ctx context.Context, tc *TwitchCommand) error
	Create(ctx context.Context, tc *TwitchCommand) error
	Delete(ctx context.Context, id string) error
	FormatCommandMessage(msg irc.ChatMessage) (string, error)
	GetCommandVariables() map[string]string
}

type TwitchCommandRepo interface {
	Fetch(ctx context.Context) (data []TwitchCommand, err error)
	FetchById(ctx context.Context, id string) (TwitchCommand, error)
	FetchByName(ctx context.Context, name string) (TwitchCommand, error)
	Update(ctx context.Context, tc *TwitchCommand) error
	Create(ctx context.Context, tc *TwitchCommand) error
	Delete(ctx context.Context, id string) error
}

type TwitchServer struct {
	API      *TwitchAPI
	PubSub   *pubsub.Client
	Commands []TwitchCommand
}

type TwitchAPI struct {
	Kraken *kraken.Client
	Helix  *helix.Client
}

type TwitchHandleCallback func(shardID int, msg irc.ChatMessage)

type TwitchCommandVariable struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TwitchCommandVariableList struct {
	Variables []TwitchCommandVariable `json:"variables"`
}

type HelixChannelDataResponseItem struct {
	ID           string `json:"broadcaster_id"`
	Login        string `json:"broadcaster_login"`
	Name         string `json:"broadcaster_name"`
	Lang         string `json:"broadcaster_language"`
	GameID       string `json:"game_id"`
	GameTopic    string `json:"game_name"`
	Title        string `json:"title"`
	ChannelDelay string `json:"delay"`
}

type HelixChannelResponse struct {
	Data []HelixChannelDataResponseItem `json:"data"`
}
