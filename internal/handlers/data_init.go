package handlers

import (
	"net/http"

	"powertab/internal/models"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

// DataInitHandler 数据初始化处理器
type DataInitHandler struct {
	initDataSvc *services.InitDataService
}

// NewDataInitHandler 创建数据初始化处理器
func NewDataInitHandler(initDataSvc *services.InitDataService) *DataInitHandler {
	return &DataInitHandler{initDataSvc: initDataSvc}
}

// InitializeUserData 初始化用户数据
// @doc "Initialize user data (categories, widgets, speed dial)"
// POST /api/data/initialize
func (h *DataInitHandler) InitializeUserData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.Response{Code: 401, Msg: "Unauthorized"})
		return
	}

	err := h.initDataSvc.InitializeUserData(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{Code: 500, Msg: "Failed to initialize user data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "User data initialized successfully"})
}

// GetAllWidgetDefinitions 获取所有小组件定义
// @doc "Get all available widget definitions"
// GET /api/data/widgets/definitions
func (h *DataInitHandler) GetAllWidgetDefinitions(c *gin.Context) {
	widgets := h.initDataSvc.GetAllWidgetDefinitions()

	type WidgetResponse struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		Description    string   `json:"description"`
		Icon           string   `json:"icon"`
		Category       string   `json:"category"`
		SupportedSizes []string `json:"supportedSizes"`
		PreviewColor   string   `json:"previewColor"`
		ComponentType  string   `json:"componentType"`
	}

	var result []WidgetResponse
	for _, w := range widgets {
		result = append(result, WidgetResponse{
			ID:             w.ID,
			Name:           w.Name,
			Description:    w.Description,
			Icon:           w.Icon,
			Category:       w.Category,
			SupportedSizes: parseSizeString(w.SupportedSizes),
			PreviewColor:   w.PreviewColor,
			ComponentType:  w.ComponentType,
		})
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: result})
}

// GetDefaultCategories 获取默认分类菜单
// @doc "Get default categories menu"
// GET /api/data/categories/default
func (h *DataInitHandler) GetDefaultCategories(c *gin.Context) {
	type CategoryResponse struct {
		Label string `json:"label"`
		Icon  string `json:"icon"`
	}

	defaultCategories := []CategoryResponse{
		{Label: "主页", Icon: "Home"},
		{Label: "AI 工具箱", Icon: "Cpu"},
		{Label: "PowerTab AI", Icon: "Sparkles"},
		{Label: "工具箱", Icon: "Grid"},
		{Label: "PDF 工具", Icon: "FileText"},
		{Label: "壁纸", Icon: "Image"},
		{Label: "日历", Icon: "Calendar"},
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: defaultCategories})
}

// GetDefaultSpeedDial 获取默认快速拨号
// @doc "Get default speed dial items"
// GET /api/data/speed-dial/default
func (h *DataInitHandler) GetDefaultSpeedDial(c *gin.Context) {
	type SpeedDialResponse struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Color string `json:"color"`
		Icon  string `json:"icon"`
	}

	defaultSpeedDial := []SpeedDialResponse{
		{Name: "Bilibili", URL: "https://www.bilibili.com", Color: "#00AEEC", Icon: "📺"},
		{Name: "GitHub", URL: "https://github.com", Color: "#181717", Icon: "💻"},
		{Name: "知乎", URL: "https://www.zhihu.com", Color: "#0084FF", Icon: "🧠"},
		{Name: "Youtube", URL: "https://www.youtube.com", Color: "#FF0000", Icon: "▶️"},
		{Name: "微博", URL: "https://weibo.com", Color: "#E6162D", Icon: "👁️"},
		{Name: "ChatGPT", URL: "https://chat.openai.com", Color: "#10A37F", Icon: "🤖"},
		{Name: "小红书", URL: "https://www.xiaohongshu.com", Color: "#FF2442", Icon: "📕"},
		{Name: "淘宝", URL: "https://taobao.com", Color: "#FF5000", Icon: "🛍️"},
	}

	c.JSON(http.StatusOK, models.Response{Code: 0, Msg: "ok", Data: defaultSpeedDial})
}

// parseSizeString 将字符串解析为大小数组
func parseSizeString(sizeStr string) []string {
	// 简单的字符串分割
	var sizes []string
	var current string

	for _, ch := range sizeStr {
		if ch == ',' {
			if current != "" {
				sizes = append(sizes, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}

	if current != "" {
		sizes = append(sizes, current)
	}

	return sizes
}
