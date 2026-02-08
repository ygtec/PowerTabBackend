package handlers

import (
	"net/http"

	"powertab/internal/models"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

type CalendarHandler struct {
	calendarSvc *services.CalendarService
}

func NewCalendarHandler(calendarSvc *services.CalendarService) *CalendarHandler {
	return &CalendarHandler{calendarSvc: calendarSvc}
}

func (h *CalendarHandler) GetEvents(c *gin.Context) {
	userID, _ := c.Get("userID")
	events, err := h.calendarSvc.GetEvents(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: events})
}

func (h *CalendarHandler) AddEvent(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.CalendarEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	event, err := h.calendarSvc.AddEvent(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: event})
}

func (h *CalendarHandler) UpdateEvent(c *gin.Context) {
	userID, _ := c.Get("userID")
	eventID := c.Param("id")
	var req models.CalendarEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	event, err := h.calendarSvc.UpdateEvent(userID.(string), eventID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: event})
}

func (h *CalendarHandler) DeleteEvent(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.calendarSvc.DeleteEvent(userID.(string), req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}
