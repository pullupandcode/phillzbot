package command

import (
	"context"
	"phillzbot/domain"
)

type discordCommandService struct {
	repo domain.DiscordRepo
}

func NewDiscordCommandService(dcr domain.DiscordRepo) domain.DiscordService {
	return &discordCommandService{
		repo: dcr,
	}
}

func (s *discordCommandService) FetchCommand(ctx context.Context) (data []domain.DiscordCommand, err error) {
	return nil, nil
}
func (s *discordCommandService) FetchCommandById(ctx context.Context, id string) (data domain.DiscordCommand, err error) {
	return data, nil
}
func (s *discordCommandService) FetchCommandByName(ctx context.Context, name string) (data domain.DiscordCommand, err error) {
	return data, nil
}
func (s *discordCommandService) UpdateCommand(ctx context.Context, tc domain.DiscordCommand) error {
	return nil
}
func (s *discordCommandService) CreateCommand(ctx context.Context, tc domain.DiscordCommand) error {
	return nil
}
func (s *discordCommandService) DeleteCommand(ctx context.Context, id string) error {
	return nil
}
