package handlers

import (
	"net/http"
	"regexp"

	"powertab/internal/models"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSvc *services.AuthService
}

func NewAuthHandler(authSvc *services.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

// isValidEmail 验证邮箱格式
func isValidEmail(email string) bool {
	// 使用简单的正则表达式验证邮箱格式
	// 该正则表达式遵循 RFC 5322 的基本要求
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	// 使用 ShouldBindJSON 但不使用 email binding 验证
	if err := c.ShouldBindJSON(&req); err != nil {
		// 检查是否是邮箱验证错误
		if err.Error() == "Key: 'RegisterRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag" {
			c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "邮箱格式不正确，请输入有效的邮箱地址"})
			return
		}
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}

	// 手动验证邮箱格式
	if req.Email == "" {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "邮箱不能为空"})
		return
	}

	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "邮箱格式不正确，请输入有效的邮箱地址（例如：user@example.com）"})
		return
	}

	// 验证密码
	if req.Password == "" {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "密码不能为空"})
		return
	}

	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "密码长度至少为 6 个字符"})
		return
	}

	// 如果未提供用户名，会在 service 中从邮箱自动生成
	user, err := h.authSvc.Register(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: user})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}

	// 手动验证邮箱格式
	if req.Email == "" {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "邮箱不能为空"})
		return
	}

	if !isValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "邮箱格式不正确，请输入有效的邮箱地址（例如：user@example.com）"})
		return
	}

	// 验证密码
	if req.Password == "" {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: "密码不能为空"})
		return
	}

	user, err := h.authSvc.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{Code: 401, Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: user})
}

func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.Response{Code: 401, Msg: "Unauthorized"})
		return
	}

	user, err := h.authSvc.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, models.Response{Code: 404, Msg: "User not found"})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: user})
}

func (h *AuthHandler) UpdateUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.Response{Code: 401, Msg: "Unauthorized"})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{Code: 400, Msg: err.Error()})
		return
	}

	user, err := h.authSvc.UpdateUser(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: user})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.Response{Code: 401, Msg: "Unauthorized"})
		return
	}

	if err := h.authSvc.Logout(userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok"})
}
