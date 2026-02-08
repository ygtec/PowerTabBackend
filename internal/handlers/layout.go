package handlers

import (
	"net/http"

	"powertab/internal/models"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

type LayoutHandler struct {
	layoutSvc *services.LayoutService
}

func NewLayoutHandler(layoutSvc *services.LayoutService) *LayoutHandler {
	return &LayoutHandler{layoutSvc: layoutSvc}
}

func (h *LayoutHandler) GetCategories(c *gin.Context) {
	userID, _ := c.Get("userID")
	cats, err := h.layoutSvc.GetCategories(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: cats})
}

func (h *LayoutHandler) SaveCategories(c *gin.Context) {
	userID, _ := c.Get("userID")
	var categories []models.Category
	if err := c.ShouldBindJSON(&categories); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.layoutSvc.SaveCategories(userID.(string), categories); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}

func (h *LayoutHandler) GetSpeedDial(c *gin.Context) {
	userID, _ := c.Get("userID")
	items, err := h.layoutSvc.GetSpeedDial(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: items})
}

func (h *LayoutHandler) SaveSpeedDial(c *gin.Context) {
	userID, _ := c.Get("userID")
	var items []models.SpeedDialItem
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.layoutSvc.SaveSpeedDial(userID.(string), items); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}

func (h *LayoutHandler) GetWidgets(c *gin.Context) {
	userID, _ := c.Get("userID")
	widgets, err := h.layoutSvc.GetWidgets(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: widgets})
}

func (h *LayoutHandler) SaveWidgets(c *gin.Context) {
	userID, _ := c.Get("userID")
	var widgets []models.Widget
	if err := c.ShouldBindJSON(&widgets); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.layoutSvc.SaveWidgets(userID.(string), widgets); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}

func (h *LayoutHandler) GetRecommendedTools(c *gin.Context) {
	userID, _ := c.Get("userID")
	tools, err := h.layoutSvc.GetRecommendedTools(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: tools})
}

func (h *LayoutHandler) SaveRecommendedTools(c *gin.Context) {
	userID, _ := c.Get("userID")
	var tools []models.RecommendedTool
	if err := c.ShouldBindJSON(&tools); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.layoutSvc.SaveRecommendedTools(userID.(string), tools); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}

func (h *LayoutHandler) GetToolOrder(c *gin.Context) {
	userID, _ := c.Get("userID")
	orders, err := h.layoutSvc.GetToolOrder(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: orders})
}

func (h *LayoutHandler) SaveToolOrder(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req struct {
		Category string `json:"category"`
		Order    []int  `json:"order"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.layoutSvc.SaveToolOrder(userID.(string), req.Category, req.Order); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}
