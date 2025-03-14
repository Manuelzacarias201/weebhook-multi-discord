package application

import (
	"context"
	"weebhook/domain"
	"weebhook/domain/entities"
)

type ReviewUseCase struct {
	repo               domain.IPayloadRepository
	discordWeebHookURL string
}

func NewReviewUseCase(repo domain.IPayloadRepository, weebHookURL string) *ReviewUseCase {
	return &ReviewUseCase{
		repo:               repo,
		discordWeebHookURL: weebHookURL,
	}
}

func (uc *ReviewUseCase) ProcessReview(ctx context.Context, payload *entities.ReviewEventPayload) error {
	discordMessage := uc.repo.FormatReviewMessage(*payload)
	err := uc.repo.SendDiscordNotification(ctx, uc.discordWeebHookURL, discordMessage)
	if err != nil {
		return err
	}
	return nil
}
