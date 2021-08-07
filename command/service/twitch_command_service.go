package command

import (
	"context"
	"phillzbot/domain"
)

type twitchCommandService struct {
	repo domain.TwitchCommandRepo
}

func NewTwitchCommandService(tcr domain.TwitchCommandRepo) domain.TwitchCommandService {
	return &twitchCommandService{
		repo: tcr,
	}
}

func (s *twitchCommandService) Fetch(ctx context.Context) (data []domain.TwitchCommand, err error) {
	return nil, nil
}
func (s *twitchCommandService) FetchById(ctx context.Context, id string) (data domain.TwitchCommand, err error) {
	return data, nil
}
func (s *twitchCommandService) FetchByName(ctx context.Context, name string) (data domain.TwitchCommand, err error) {
	return data, nil
}
func (s *twitchCommandService) Update(ctx context.Context, tc *domain.TwitchCommand) error {
	return nil
}
func (s *twitchCommandService) Create(ctx context.Context, tc *domain.TwitchCommand) error {
	return nil
}
func (s *twitchCommandService) Delete(ctx context.Context, id string) error {
	return nil
}
