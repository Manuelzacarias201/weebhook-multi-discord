package routes

import (
	"weebhook/infraestructure/controller"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, webhookHandler *controller.WebhookHandler, reviewHandler *controller.ReviewHandler) {
	routes := router.Group("pull_request")
	{
		routes.POST("/webhook", webhookHandler.HandlePullRequest)
	}

	reviews := router.Group("review")
	{
		reviews.POST("/webhook", reviewHandler.HandleReview)
	}
}
