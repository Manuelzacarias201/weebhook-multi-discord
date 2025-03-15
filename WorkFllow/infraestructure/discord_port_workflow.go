package infraestructure

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)


func post_discord_actions(msg string) int {
	discord_webhook_url := os.Getenv("DISCORD_WEBHOOK_URL_ACTIONS") 
	if discord_webhook_url == "" {
		log.Println("Error: Webhook de Actions no está configurado")
		return 500
	}

	payload := map[string]string{
		"content": msg,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error al serializar JSON para Discord Actions")
		return 500
	}

	resp, err := http.Post(discord_webhook_url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error al enviar mensaje a Discord Actions")
		return 500
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		log.Printf("Error al enviar mensaje de Actions, código: %d", resp.StatusCode)
		return 400
	}

	return 200
}
