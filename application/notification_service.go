package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"pull_request_webhook/domain"
)

type NotificationService struct {
	settings *domain.NotificationSettings
	client   *http.Client
}

type NotificationPayload struct {
	Content   string `json:"content"`
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

func NewNotificationService(settings *domain.NotificationSettings) *NotificationService {
	return &NotificationService{
		settings: settings,
		client: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

func (s *NotificationService) Notify(notificationType domain.NotificationType, message string) error {
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

	payload := NotificationPayload{
		Content:   message,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return s.sendNotification(config.WebhookURL, payload)
}

func (s *NotificationService) sendNotification(webhookURL string, payload NotificationPayload) error {
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
