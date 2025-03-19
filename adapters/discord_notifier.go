package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"pull_request_webhook/domain"
	"pull_request_webhook/ports"
)

type DiscordNotifier struct {
	settings *domain.NotificationSettings
	client   *http.Client
}

type DiscordPayload struct {
	Content   string `json:"content"`
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

func NewDiscordNotifier(settings *domain.NotificationSettings) ports.NotificationService {
	return &DiscordNotifier{
		settings: settings,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (s *DiscordNotifier) SendDevelopmentNotification(message string) error {
	return s.Notify(domain.DevNotification, message)
}

func (s *DiscordNotifier) SendTestingNotification(message string) error {
	return s.Notify(domain.TestNotification, message)
}

func (s *DiscordNotifier) SendGeneralNotification(message string) error {
	return s.Notify(domain.GeneralNotification, message)
}

func (s *DiscordNotifier) Notify(notificationType domain.NotificationType, message string) error {
	var config domain.NotificationConfig
	switch notificationType {
	case domain.DevNotification:
		config = s.settings.Dev
	case domain.TestNotification:
		config = s.settings.Test
	case domain.GeneralNotification:
		config = s.settings.General
	default:
		return fmt.Errorf("tipo de notificación no válido")
	}

	if !config.Enabled {
		return fmt.Errorf("canal de notificación deshabilitado")
	}

	payload := DiscordPayload{
		Content:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return s.sendNotification(config.WebhookURL, payload)
}

func (s *DiscordNotifier) sendNotification(webhookURL string, payload DiscordPayload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al preparar la notificación: %v", err)
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar la notificación: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error en la respuesta del servidor: %d", resp.StatusCode)
	}

	return nil
}
