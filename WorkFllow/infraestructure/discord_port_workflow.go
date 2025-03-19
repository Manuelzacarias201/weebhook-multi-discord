package infraestructure

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"pull_request_webhook/domain"
)

func post_discord_actions(msg string, notificationType domain.NotificationType) int {
	var webhookURL string
	switch notificationType {
	case domain.DevNotification:
		webhookURL = os.Getenv("DISCORD_WEBHOOK_URL_DEV")
	case domain.TestNotification:
		webhookURL = os.Getenv("DISCORD_WEBHOOK_URL_ACTIONS")
	case domain.GeneralNotification:
		webhookURL = os.Getenv("DISCORD_WEBHOOK_URL_GENERAL")
	default:
		log.Println("Error: Tipo de notificación no válido")
		return 500
	}

	if webhookURL == "" {
		log.Printf("Error: Webhook para el tipo %d no está configurado", notificationType)
		return 500
	}

	payload := map[string]string{
		"content": msg,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error al serializar JSON para Discord")
		return 500
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error al enviar mensaje a Discord")
		return 500
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Printf("Error al enviar mensaje, código: %d", resp.StatusCode)
		return 400
	}

	return 200
}
