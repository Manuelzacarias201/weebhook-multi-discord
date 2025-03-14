package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"weebhook/domain/entities"
)

type PayloadRepository struct{}

func NewPayloadRepository() *PayloadRepository {
	return &PayloadRepository{}
}

func (r *PayloadRepository) ProcessPullRequestPayload(ctx context.Context, payload entities.PullRequestEventPayload) error {
	return nil
}

func (r *PayloadRepository) FormatDiscordMessage(payload entities.PullRequestEventPayload) interface{} {
	// Mapeo entre acciones de PR y colores para los embeds
	colorMap := map[string]int{
		"opened":      5025616,  // Verde
		"closed":      15158332, // Rojo
		"reopened":    3447003,  // Azul
		"synchronize": 16776960, // Amarillo
	}

	// Determinar color basado en la acción
	color, ok := colorMap[payload.Action]
	if !ok {
		color = 9807270 // Gris por defecto para acciones no mapeadas
	}

	// Crear el mensaje con formato de embeds para Discord
	return map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title": fmt.Sprintf("Pull Request: %s", payload.PullRequest.Title),
				"url":   payload.PullRequest.URL,
				"color": color,
				"author": map[string]interface{}{
					"name":     payload.PullRequest.User.Login,
					"icon_url": payload.PullRequest.User.URL,
				},
				"fields": []map[string]interface{}{
					{
						"name":   "Repository",
						"value":  payload.Repository.FullName,
						"inline": true,
					},
					{
						"name":   "Action",
						"value":  payload.Action,
						"inline": true,
					},
					{
						"name": "Branch",
						"value": fmt.Sprintf("%s → %s",
							payload.PullRequest.Head.Ref,
							payload.PullRequest.Base.Ref),
						"inline": false,
					},
				},
			},
		},
	}
}

func (r *PayloadRepository) SendDiscordNotification(ctx context.Context, webhookURL string, message interface{}) error {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling discord message: %w", err)
	}

	// Imprime el payload para depuración
	fmt.Printf("Sending to Discord: %s\n", string(jsonData))

	// Verificar que no estamos enviando nil o un mensaje vacío
	if message == nil {
		return fmt.Errorf("cannot send nil message to Discord")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending discord notification: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Leer el cuerpo de la respuesta para obtener más detalles
		responseBody := new(bytes.Buffer)
		_, _ = responseBody.ReadFrom(resp.Body)

		return fmt.Errorf("discord API responded with status: %d, body: %s",
			resp.StatusCode, responseBody.String())
	}

	return nil
}

func (r *PayloadRepository) ProcessReviewPayload(ctx context.Context, payload entities.ReviewEventPayload) error {
	return nil
}

func (r *PayloadRepository) FormatReviewMessage(payload entities.ReviewEventPayload) interface{} {
	return map[string]interface{}{
		"embeds": []map[string]interface{}{
			{
				"title":       fmt.Sprintf("Review #%d: %s", payload.Review.ID, payload.Review.State),
				"description": payload.Review.Body,
				"url":         payload.PullRequest.URL,
				"color":       3447003,
				"author": map[string]interface{}{
					"name":     payload.Review.User.Login,
					"icon_url": payload.Review.User.URL,
				},
				"fields": []map[string]interface{}{
					{
						"name":   "Repository",
						"value":  payload.Repository.FullName,
						"inline": true,
					},
					{
						"name":   "Action",
						"value":  payload.Action,
						"inline": true,
					},
					{
						"name":   "State",
						"value":  payload.Review.State,
						"inline": true,
					},
				},
			},
		},
	}
}
