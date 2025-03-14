package application

import (
	"context"
	"weebhook/domain"
	"weebhook/domain/entities"
)

type PayloadUseCase struct {
	repo               domain.IPayloadRepository
	discordWeebHookURL string
}

func NewPayloadUseCase(repo domain.IPayloadRepository, weebHookURL string) *PayloadUseCase {
	return &PayloadUseCase{
		repo:               repo,
		discordWeebHookURL: weebHookURL,
	}
}

func (uc *PayloadUseCase) ProcessPullRequest(ctx context.Context, payload *entities.PullRequestEventPayload) error {
	discordMessage := uc.repo.FormatDiscordMessage(*payload)
	err := uc.repo.SendDiscordNotification(ctx, uc.discordWeebHookURL, discordMessage)
	if err != nil {
		return err
	}
	return nil
}
