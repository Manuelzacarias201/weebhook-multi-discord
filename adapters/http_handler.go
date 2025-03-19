package adapters

import (
	"encoding/json"
	"net/http"

	"pull_request_webhook/application"
	"pull_request_webhook/domain"
)

type WebhookHandler struct {
	eventProcessor *application.EventProcessor
}

func NewWebhookHandler(eventProcessor *application.EventProcessor) *WebhookHandler {
	return &WebhookHandler{
		eventProcessor: eventProcessor,
	}
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var event domain.GitHubEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Error al procesar el evento", http.StatusBadRequest)
		return
	}

	eventType := r.Header.Get("X-GitHub-Event")
	switch eventType {
	case "pull_request":
		h.eventProcessor.ProcessPullRequestEvent(event)
	case "workflow_run":
		h.eventProcessor.ProcessWorkflowEvent(event)
	}

	w.WriteHeader(http.StatusOK)
}
