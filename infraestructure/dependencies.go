package infraestructure

import (
	"os"
	"weebhook/application"
	"weebhook/infraestructure/controller"
	"weebhook/infraestructure/repositories"

	"github.com/joho/godotenv"
)

func Init() (*controller.WebhookHandler, *controller.ReviewHandler) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	payloadRepo := repositories.NewPayloadRepository()

	discordWebhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	payloadUseCase := application.NewPayloadUseCase(payloadRepo, discordWebhookURL)
	reviewUseCase := application.NewReviewUseCase(payloadRepo, discordWebhookURL)

	webhookHandler := controller.NewWebhookHandler(*payloadUseCase)
	reviewHandler := controller.NewReviewHandler(*reviewUseCase)

	return webhookHandler, reviewHandler
}
