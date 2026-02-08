package handlers

import (
	"net/http"

	"powertab/internal/models"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	aiSvc *services.AIService
}

func NewAIHandler(aiSvc *services.AIService) *AIHandler {
	return &AIHandler{aiSvc: aiSvc}
}

func (h *AIHandler) GetSessions(c *gin.Context) {
	userID, _ := c.Get("userID")
	sessions, err := h.aiSvc.GetSessions(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: sessions})
}

func (h *AIHandler) SaveSessions(c *gin.Context) {
	userID, _ := c.Get("userID")
	var sessions []models.ChatSession
	if err := c.ShouldBindJSON(&sessions); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.aiSvc.SaveSessions(userID.(string), sessions); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}
