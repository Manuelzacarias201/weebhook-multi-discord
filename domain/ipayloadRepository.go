package domain

import (
	"context"
	"weebhook/domain/entities"
)

type IPayloadRepository interface {
	ProcessPullRequestPayload(ctx context.Context, payload entities.PullRequestEventPayload) error
	SendDiscordNotification(ctx context.Context, weebHookURL string, message interface{}) error
	FormatDiscordMessage(payload entities.PullRequestEventPayload) interface{}
	ProcessReviewPayload(ctx context.Context, payload entities.ReviewEventPayload) error
	FormatReviewMessage(payload entities.ReviewEventPayload) interface{}
}
