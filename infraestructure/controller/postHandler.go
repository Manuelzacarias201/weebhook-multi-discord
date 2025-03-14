package controller

import (
	"fmt"
	"net/http"
	"weebhook/application"
	"weebhook/domain/entities"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	payloadUseCase application.PayloadUseCase
}

func NewWebhookHandler(payloadUseCase application.PayloadUseCase) *WebhookHandler {
	return &WebhookHandler{payloadUseCase: payloadUseCase}
}

func (h *WebhookHandler) HandlePullRequest(g *gin.Context) {
	var payload entities.PullRequestEventPayload

	if err := g.ShouldBindJSON(&payload); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload", "details": err.Error()})
		fmt.Printf("Error binding JSON: %v\n", err)
		return
	}

	err := h.payloadUseCase.ProcessPullRequest(g.Request.Context(), &payload)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		fmt.Printf("Error processing pull request: %v\n", err)
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "OK"})
	fmt.Println("Pull request handled successfully")
}
