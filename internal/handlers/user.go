package handlers

import (
	"net/http"
	"powertab/internal/models"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc *services.UserService
}

func NewUserHandler(userSvc *services.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

func (h *UserHandler) GetPreferences(c *gin.Context) {
	userID, _ := c.Get("userID")
	pref, err := h.userSvc.GetPreferences(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: pref})
}

func (h *UserHandler) UpdatePreferences(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.PreferencesUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	pref, err := h.userSvc.UpdatePreferences(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: pref})
}

func (h *UserHandler) ListSearchEngines(c *gin.Context) {
	userID, _ := c.Get("userID")
	engines, err := h.userSvc.ListSearchEngines(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: engines})
}

func (h *UserHandler) AddSearchEngine(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.SearchEngineRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	engine, err := h.userSvc.AddSearchEngine(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: engine})
}

func (h *UserHandler) DeleteSearchEngine(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.userSvc.DeleteSearchEngine(userID.(string), req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}

func (h *UserHandler) ListWallpapers(c *gin.Context) {
	userID, _ := c.Get("userID")
	wallpapers, err := h.userSvc.ListWallpapers(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: wallpapers})
}

func (h *UserHandler) AddWallpaper(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req models.WallpaperRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	wallpaper, err := h.userSvc.AddWallpaper(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: wallpaper})
}

func (h *UserHandler) DeleteWallpaper(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req struct {
		ID string `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}
	if err := h.userSvc.DeleteWallpaper(userID.(string), req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}
