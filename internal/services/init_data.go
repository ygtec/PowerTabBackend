package services

import (
	"powertab/config"
	"powertab/internal/models"
	"powertab/internal/utils"

	"gorm.io/gorm"
)

// InitDataService 初始化数据服务
type InitDataService struct {
	db *gorm.DB
}

// NewInitDataService 创建初始化数据服务
func NewInitDataService() *InitDataService {
	return &InitDataService{db: config.GetDB()}
}

// InitializeDefaultCategories 初始化默认分类菜单
func (s *InitDataService) InitializeDefaultCategories(userID string) error {
	defaultCategories := []models.Category{
		{
			ID:        "dashboard",
			UserID:    userID,
			Label:     "主页",
			Icon:      "Home",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        "ai-tools",
			UserID:    userID,
			Label:     "AI 工具箱",
			Icon:      "Cpu",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        "powertab-ai",
			UserID:    userID,
			Label:     "PowerTab AI",
			Icon:      "Sparkles",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        "productivity",
			UserID:    userID,
			Label:     "工具箱",
			Icon:      "Grid",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        "pdf-tools",
			UserID:    userID,
			Label:     "PDF 工具",
			Icon:      "FileText",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        "wallpapers",
			UserID:    userID,
			Label:     "壁纸",
			Icon:      "Image",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        "calendar",
			UserID:    userID,
			Label:     "日历",
			Icon:      "Calendar",
			CreatedAt: models.GetCurrentTime(),
		},
	}

	// 检查是否已存在该用户的分类
	var count int64
	s.db.Model(&models.Category{}).Where("user_id = ?", userID).Count(&count)

	// 如果还没有分类，则插入默认分类
	if count == 0 {
		return s.db.CreateInBatches(defaultCategories, 100).Error
	}

	return nil
}

// InitializeDefaultWidgets 初始化默认小组件
func (s *InitDataService) InitializeDefaultWidgets(userID string) error {
	defaultWidgets := []models.Widget{
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			WidgetID:  "weather",
			PageID:    "dashboard",
			Size:      "medium",
			Settings:  []byte(`{"location":"Beijing"}`),
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			WidgetID:  "clock",
			PageID:    "dashboard",
			Size:      "small",
			Settings:  []byte(`{"format":"24h"}`),
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			WidgetID:  "todo",
			PageID:    "dashboard",
			Size:      "medium",
			Settings:  []byte(`{}`),
			CreatedAt: models.GetCurrentTime(),
		},
	}

	// 检查是否已存在该用户的小组件
	var count int64
	s.db.Model(&models.Widget{}).Where("user_id = ?", userID).Count(&count)

	// 如果还没有小组件，则插入默认小组件
	if count == 0 {
		return s.db.CreateInBatches(defaultWidgets, 100).Error
	}

	return nil
}

// InitializeDefaultSpeedDial 初始化默认快速拨号
func (s *InitDataService) InitializeDefaultSpeedDial(userID string) error {
	defaultSpeedDials := []models.SpeedDialItem{
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "Bilibili",
			URL:       "https://www.bilibili.com",
			Color:     "#00AEEC",
			Icon:      "📺",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "GitHub",
			URL:       "https://github.com",
			Color:     "#181717",
			Icon:      "💻",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "知乎",
			URL:       "https://www.zhihu.com",
			Color:     "#0084FF",
			Icon:      "🧠",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "Youtube",
			URL:       "https://www.youtube.com",
			Color:     "#FF0000",
			Icon:      "▶️",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "微博",
			URL:       "https://weibo.com",
			Color:     "#E6162D",
			Icon:      "👁️",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "ChatGPT",
			URL:       "https://chat.openai.com",
			Color:     "#10A37F",
			Icon:      "🤖",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "小红书",
			URL:       "https://www.xiaohongshu.com",
			Color:     "#FF2442",
			Icon:      "📕",
			CreatedAt: models.GetCurrentTime(),
		},
		{
			ID:        utils.GenerateID(),
			UserID:    userID,
			Name:      "淘宝",
			URL:       "https://taobao.com",
			Color:     "#FF5000",
			Icon:      "🛍️",
			CreatedAt: models.GetCurrentTime(),
		},
	}

	// 检查是否已存在该用户的快速拨号
	var count int64
	s.db.Model(&models.SpeedDialItem{}).Where("user_id = ?", userID).Count(&count)

	// 如果还没有快速拨号，则插入默认快速拨号
	if count == 0 {
		return s.db.CreateInBatches(defaultSpeedDials, 100).Error
	}

	return nil
}

// InitializeUserData 一次性初始化用户的所有数据
func (s *InitDataService) InitializeUserData(userID string) error {
	// 初始化分类
	if err := s.InitializeDefaultCategories(userID); err != nil {
		return err
	}

	// 初始化小组件
	if err := s.InitializeDefaultWidgets(userID); err != nil {
		return err
	}

	// 初始化快速拨号
	if err := s.InitializeDefaultSpeedDial(userID); err != nil {
		return err
	}

	return nil
}

// GetAllWidgetDefinitions 获取所有小组件定义
func (s *InitDataService) GetAllWidgetDefinitions() []models.WidgetDefinition {
	return []models.WidgetDefinition{
		{
			ID:             "weather",
			Name:           "天气",
			Description:    "实时查看当地天气预报，了解温度、降水、风力等信息。",
			Icon:           "🌤️",
			Category:       "information",
			SupportedSizes: "small,medium,large",
			PreviewColor:   "from-blue-400 to-cyan-500",
			ComponentType:  "weather",
		},
		{
			ID:             "clock",
			Name:           "时钟",
			Description:    "显示当前时间，支持12小时和24小时制。",
			Icon:           "🕐",
			Category:       "information",
			SupportedSizes: "small,medium",
			PreviewColor:   "from-purple-400 to-purple-600",
			ComponentType:  "clock",
		},
		{
			ID:             "todo",
			Name:           "待办事项",
			Description:    "列出待办事项，帮助用户管理和追踪任务。",
			Icon:           "✅",
			Category:       "efficiency",
			SupportedSizes: "small,medium,large",
			PreviewColor:   "from-green-400 to-emerald-600",
			ComponentType:  "todo",
		},
		{
			ID:             "news",
			Name:           "新闻",
			Description:    "实时获取最新新闻，掌握时事动态。",
			Icon:           "📰",
			Category:       "information",
			SupportedSizes: "medium,large",
			PreviewColor:   "from-orange-400 to-red-600",
			ComponentType:  "news",
		},
		{
			ID:             "stock",
			Name:           "股票",
			Description:    "让您轻松关注股票和股市的动态，沪深、港股、美股全球市场实时行情。",
			Icon:           "📈",
			Category:       "finance",
			SupportedSizes: "small,medium,large",
			PreviewColor:   "from-green-500 to-emerald-700",
			ComponentType:  "stock",
		},
		{
			ID:             "nba",
			Name:           "NBA赛事",
			Description:    "NBA赛程、比分、排名，实时掌握比赛动态。",
			Icon:           "🏀",
			Category:       "sports",
			SupportedSizes: "small,medium,large",
			PreviewColor:   "from-blue-600 to-indigo-800",
			ComponentType:  "nba",
		},
		{
			ID:             "pomodoro",
			Name:           "番茄时钟",
			Description:    "通过设定专注工作和休息间隔，帮助用户保持专注，减少疲劳。",
			Icon:           "🍅",
			Category:       "efficiency",
			SupportedSizes: "small,medium",
			PreviewColor:   "from-red-400 to-orange-500",
			ComponentType:  "pomodoro",
		},
		{
			ID:             "translator",
			Name:           "翻译",
			Description:    "快速文本翻译，支持多种语言互译。",
			Icon:           "🌐",
			Category:       "tool",
			SupportedSizes: "medium,large",
			PreviewColor:   "from-purple-400 to-purple-600",
			ComponentType:  "translator",
		},
	}
}
