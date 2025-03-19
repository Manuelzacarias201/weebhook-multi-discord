package application

import (
	"encoding/json"
	"log"
	"pull_request_webhook/domain/value_objects"
)

func ProcessWorkflowEvent(rawData []byte) string {
	var eventPayload value_objects.WorkflowRunEvent

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Println("Error al deserializar payload de Actions")
		return "ERROR"
	}

	log.Printf("Evento de GitHub Actions recibido: %s", eventPayload.Action)
	log.Printf("Workflow Name: %s", eventPayload.WorkflowRun.Name)
	log.Printf("Conclusion: %s", eventPayload.WorkflowRun.Conclusion)
	log.Printf("Repository: %s", eventPayload.Repository.FullName)
	log.Printf("Workflow URL: %s", eventPayload.WorkflowRun.HTMLURL)

	if eventPayload.WorkflowRun.Name == "" {
		eventPayload.WorkflowRun.Name = "Desconocido"
	}
	if eventPayload.WorkflowRun.Conclusion == "" {
		eventPayload.WorkflowRun.Conclusion = "Desconocido"
	}

	return GenerateMessageToDiscordForActions(
		eventPayload.Action,
		eventPayload.WorkflowRun.Name,
		eventPayload.WorkflowRun.Conclusion,
		eventPayload.Repository.FullName,
		eventPayload.WorkflowRun.HTMLURL,
	)
}
