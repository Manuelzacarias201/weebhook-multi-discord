package controller

import (
	"net/http"
	"weebhook/application"
	"weebhook/domain/entities"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	reviewUseCase application.ReviewUseCase
}

func NewReviewHandler(reviewUseCase application.ReviewUseCase) *ReviewHandler {
	return &ReviewHandler{reviewUseCase: reviewUseCase}
}

func (h *ReviewHandler) HandleReview(g *gin.Context) {
	var payload entities.ReviewEventPayload

	if err := g.ShouldBindJSON(&payload); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload", "details": err.Error()})
		return
	}

	err := h.reviewUseCase.ProcessReview(g.Request.Context(), &payload)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "details": err.Error()})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "OK"})
}
