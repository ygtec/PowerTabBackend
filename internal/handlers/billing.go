package handlers

import (
	"net/http"

	"powertab/internal/models"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

type BillingHandler struct {
	billingSvc *services.BillingService
}

func NewBillingHandler(billingSvc *services.BillingService) *BillingHandler {
	return &BillingHandler{billingSvc: billingSvc}
}

func (h *BillingHandler) ConsumePoints(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.ConsumPointsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	user, err := h.billingSvc.ConsumePoints(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "Insufficient points"})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: user})
}

func (h *BillingHandler) Subscribe(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.SubscribePlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	user, err := h.billingSvc.SubscribePlan(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: user})
}

func (h *BillingHandler) BuyCredits(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.BuyCreditsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	user, err := h.billingSvc.BuyCredits(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: user})
}
