package application

import (
	"encoding/json"
	"log"
	"pull_request_webhook/domain/value_objects"
)

func ProcessPullRequestEvent(rawData []byte) string {
	var eventPayload value_objects.PullRequestEvent

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Println("Error al serializar payload")
		return "ERROR"
	}

	log.Printf("Evento pull request recibido con acción de %s", eventPayload.Action)

	base := eventPayload.PullRequest.Base.Ref
	titulo := eventPayload.PullRequest.Title
	repoFullName := eventPayload.Repository.FullName
	user := eventPayload.PullRequest.User.Login
	urlPullRequest := eventPayload.PullRequest.HTMLURL
	action := eventPayload.Action
	merged := eventPayload.PullRequest.Merged

	return GenerateMessageToDiscordForPullRequest(action, base, titulo, repoFullName, user, urlPullRequest, merged)
}

func GenerateMessageToDiscordForPullRequest(action, base, titulo, repoFullName, user, urlPullRequest string, merged bool) string {
	var message string
	switch action {
	case "opened":
		message = "🔵 Nuevo Pull Request creado"
	case "reopened":
		message = "🔄 Pull Request reabierto"
	case "ready_for_review":
		message = "👀 Pull Request listo para revisión"
	case "closed":
		if merged {
			message = "✅ Pull Request fusionado exitosamente"
		} else {
			message = "❌ Pull Request cerrado sin fusionar"
		}
	default:
		message = "📝 Actualización de Pull Request"
	}

	return message + "\n" +
		"**Repositorio:** " + repoFullName + "\n" +
		"**Título:** " + titulo + "\n" +
		"**Autor:** " + user + "\n" +
		"**Base:** " + base + "\n" +
		"**URL:** " + urlPullRequest
}
