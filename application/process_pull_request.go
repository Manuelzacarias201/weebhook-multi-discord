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

	return GenerateMessageToDiscord(action, base, titulo, repoFullName, user, urlPullRequest, merged)
}
