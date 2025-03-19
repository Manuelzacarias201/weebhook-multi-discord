package adapters

import (
	"pull_request_webhook/application"
	"pull_request_webhook/domain"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	eventProcessor *application.EventProcessor
}

func NewGinHandler(eventProcessor *application.EventProcessor) *GinHandler {
	return &GinHandler{
		eventProcessor: eventProcessor,
	}
}

func (h *GinHandler) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/github/webhook", func(c *gin.Context) {
			var event domain.GitHubEvent
			if err := c.ShouldBindJSON(&event); err != nil {
				c.JSON(400, gin.H{"error": "Error al procesar el evento"})
				return
			}

			eventType := c.GetHeader("X-GitHub-Event")
			switch eventType {
			case "pull_request":
				h.eventProcessor.ProcessPullRequestEvent(event)
			case "workflow_run":
				h.eventProcessor.ProcessWorkflowEvent(event)
			}

			c.Status(200)
		})
	}
}
