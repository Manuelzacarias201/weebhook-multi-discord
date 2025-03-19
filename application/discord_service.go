package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"pull_request_webhook/domain"
)

type DiscordService struct {
	webhookURLs *domain.DiscordChannels
}

type DiscordMessage struct {
	Content string `json:"content"`
}

func NewDiscordService(channels *domain.DiscordChannels) *DiscordService {
	return &DiscordService{
		webhookURLs: channels,
	}
}

func (s *DiscordService) SendDevelopmentNotification(message string) error {
	return s.sendMessage(s.webhookURLs.DevelopmentChannel, message)
}

func (s *DiscordService) SendTestingNotification(message string) error {
	return s.sendMessage(s.webhookURLs.TestingChannel, message)
}

func (s *DiscordService) SendGeneralNotification(message string) error {
	return s.sendMessage(s.webhookURLs.GeneralChannel, message)
}

func (s *DiscordService) sendMessage(webhookURL, message string) error {
	payload := DiscordMessage{
		Content: message,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al preparar el mensaje: %v", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error al enviar el mensaje: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error en la respuesta de Discord: %d", resp.StatusCode)
	}

	return nil
}
