package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID         string     `gorm:"type:varchar(191);primaryKey" json:"id" comment:"用户ID"`
	Username   string     `gorm:"type:varchar(100);uniqueIndex:idx_users_username,length:100;not null" json:"username" comment:"用户名"`
	Email      string     `gorm:"type:varchar(100);uniqueIndex:idx_users_email,length:100;not null" json:"email" comment:"邮箱"`
	Password   string     `gorm:"type:varchar(255);not null" json:"-" comment:"密码哈希"`
	Avatar     string     `gorm:"type:varchar(512)" json:"avatar" comment:"头像URL"`
	PlanLevel  string     `gorm:"type:varchar(32);default:'free'" json:"plan_level" comment:"订阅计划等级:free/pro/ultra"`
	PlanExpiry *time.Time `gorm:"type:datetime" json:"plan_expiry" comment:"计划过期时间"`
	Points     int        `gorm:"default:0" json:"points" comment:"用户积分"`
	Token      string     `gorm:"type:varchar(512)" json:"token" comment:"JWT令牌"`
	CreatedAt  time.Time  `json:"created_at" comment:"创建时间"`
	UpdatedAt  time.Time  `json:"updated_at" comment:"更新时间"`
}

// UserPreference 用户偏好设置
type UserPreference struct {
	ID            string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"偏好设置ID"`
	UserID        string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Theme         string    `gorm:"type:varchar(32);default:'auto'" json:"theme" comment:"主题:auto/light/dark"`
	ShowSeconds   bool      `gorm:"default:true" json:"show_seconds" comment:"是否显示秒"`
	Notifications bool      `gorm:"default:true" json:"notifications" comment:"是否启用通知"`
	Language      string    `gorm:"type:varchar(32);default:'zh'" json:"language" comment:"语言设置"`
	DefaultEngine string    `gorm:"type:varchar(64);default:'baidu'" json:"default_engine" comment:"默认搜索引擎"`
	WallpaperURL  string    `gorm:"type:varchar(512)" json:"wallpaper_url" comment:"壁纸URL"`
	AutoWallpaper bool      `gorm:"default:false" json:"auto_wallpaper" comment:"是否自动更换壁纸"`
	IconSize      string    `gorm:"type:varchar(32);default:'medium'" json:"icon_size" comment:"图标大小:small/medium/large"`
	UpdatedAt     time.Time `json:"updated_at" comment:"更新时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// SearchEngine 搜索引擎
type SearchEngine struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"搜索引擎ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Name      string    `gorm:"type:varchar(128);not null" json:"name" comment:"引擎名称"`
	URL       string    `gorm:"type:varchar(512);not null" json:"url" comment:"搜索URL模板"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// SavedWallpaper 保存的壁纸
type SavedWallpaper struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"壁纸ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	URL       string    `gorm:"type:varchar(512);not null" json:"url" comment:"壁纸URL"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// Category 分类
type Category struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"分类ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Label     string    `gorm:"type:varchar(128);not null" json:"label" comment:"分类名称"`
	Icon      string    `gorm:"type:varchar(256);not null" json:"icon" comment:"分类图标"`
	Order     int       `json:"order" comment:"排序顺序"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// SpeedDialItem 快速拨号项
type SpeedDialItem struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"一级按餐ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Name      string    `gorm:"type:varchar(128);not null" json:"name" comment:"名称"`
	URL       string    `gorm:"type:varchar(512);not null" json:"url" comment:"网址"`
	Color     string    `gorm:"type:varchar(32)" json:"color" comment:"颜色"`
	Icon      string    `gorm:"type:varchar(256)" json:"icon" comment:"图标"`
	Order     int       `json:"order" comment:"排序"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// Widget 小工具
type Widget struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"小工具ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	WidgetID  string    `gorm:"type:varchar(191);not null" json:"widget_id" comment:"小工具类ID"`
	PageID    string    `gorm:"type:varchar(191)" json:"page_id" comment:"页面ID"`
	Size      string    `gorm:"type:varchar(32)" json:"size" comment:"大小:small/medium/large"`
	Settings  []byte    `gorm:"type:longblob" json:"settings" comment:"配置 JSON"`
	Order     int       `json:"order" comment:"排序"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// RecommendedTool 推荐工具
type RecommendedTool struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"推荐工具ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	ToolID    int       `gorm:"not null" json:"tool_id" comment:"工具ID"`
	Order     int       `json:"order" comment:"排序"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// ToolOrder 工具排序
type ToolOrder struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"排序ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Category  string    `gorm:"type:varchar(128);not null;index:idx_user_category,composite:yes" json:"category" comment:"分类"`
	Orders    []byte    `gorm:"type:longblob" json:"orders" comment:"排序 JSON"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" comment:"更新时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// CalendarEvent 日历事件
type CalendarEvent struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"事件ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Title     string    `gorm:"type:varchar(256);not null" json:"title" comment:"事件标简"`
	Date      string    `gorm:"type:varchar(32);not null;index:idx_user_date,composite:yes" json:"date" comment:"日期 YYYY-MM-DD"`
	Time      string    `gorm:"type:varchar(32)" json:"time" comment:"時間 HH:mm"`
	Type      string    `gorm:"type:varchar(32);default:'todo'" json:"type" comment:"类型:todo/holiday/birthday/anniversary/countdown"`
	Completed bool      `gorm:"default:false" json:"completed" comment:"是否完成"`
	Recurring bool      `gorm:"default:false" json:"recurring" comment:"是否重复"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" comment:"更新时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// ChatMessage 聊天消息
type ChatMessage struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"消息ID"`
	SessionID string    `gorm:"type:varchar(191);not null;index" json:"session_id" comment:"会话ID"`
	Role      string    `gorm:"type:varchar(32);not null" json:"role" comment:"角色:user/model"`
	Content   string    `gorm:"type:longtext" json:"content" comment:"消息内容"`
	Type      string    `gorm:"type:varchar(32);default:'text'" json:"type" comment:"类型:text/image/video"`
	MediaURL  string    `gorm:"type:varchar(512)" json:"media_url" comment:"媒体URL"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	Session ChatSession `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"-"`
}

// ChatSession 聊天会话
type ChatSession struct {
	ID        string        `gorm:"type:varchar(191);primaryKey" json:"id" comment:"会话 ID"`
	UserID    string        `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Title     string        `gorm:"type:varchar(256);not null" json:"title" comment:"会话标题"`
	UpdatedAt time.Time     `json:"updated_at" comment:"更新时间"`
	Messages  []ChatMessage `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"messages" comment:"消息列表"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// PointConsumption 积分消費记录
type PointConsumption struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"消費ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	Action    string    `gorm:"type:varchar(128);not null" json:"action" comment:"操作:chat_std/chat_pro/image_gen/etc"`
	Points    int       `gorm:"not null" json:"points" comment:"消費积分"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// ToolUsageLog 工具使用日志
type ToolUsageLog struct {
	ID        string    `gorm:"type:varchar(191);primaryKey" json:"id" comment:"日志ID"`
	UserID    string    `gorm:"type:varchar(191);not null;index" json:"user_id" comment:"用户ID"`
	ToolID    string    `gorm:"type:varchar(191);not null;index" json:"tool_id" comment:"工具ID"`
	Action    string    `gorm:"type:varchar(128)" json:"action" comment:"操作"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}

// BeforeCreate GORM 钩子
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = u.generateID()
	}
	return nil
}

func (u *User) generateID() string {
	return "user_" + time.Now().Format("20060102150405")
}

// Response 通用 API 响应
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username"` // 可选，如果不提供则从 email 自动生成
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// PreferencesUpdateRequest 偏好设置更新请求
type PreferencesUpdateRequest struct {
	Theme         string `json:"theme"`
	ShowSeconds   bool   `json:"show_seconds"`
	Notifications bool   `json:"notifications"`
	Language      string `json:"language"`
	DefaultEngine string `json:"default_engine"`
	WallpaperURL  string `json:"wallpaper_url"`
	AutoWallpaper bool   `json:"auto_wallpaper"`
	IconSize      string `json:"icon_size"`
}

// SearchEngineRequest 搜索引擎请求
type SearchEngineRequest struct {
	Name string `json:"name" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

// WallpaperRequest 壁纸请求
type WallpaperRequest struct {
	URL string `json:"url" binding:"required"`
}

// CalendarEventRequest 日历事件请求
type CalendarEventRequest struct {
	Title     string `json:"title" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Time      string `json:"time"`
	Type      string `json:"type"`
	Completed bool   `json:"completed"`
	Recurring bool   `json:"recurring"`
}

// CategoryRequest 分类请求
type CategoryRequest struct {
	Label string `json:"label" binding:"required"`
	Icon  string `json:"icon" binding:"required"`
}

// SpeedDialRequest 快速拨号请求
type SpeedDialRequest struct {
	Name  string `json:"name" binding:"required"`
	URL   string `json:"url" binding:"required"`
	Color string `json:"color"`
	Icon  string `json:"icon"`
}

// AddChatMessageRequest 添加聊天消息请求
type AddChatMessageRequest struct {
	SessionID string `json:"session_id" binding:"required"`
	Role      string `json:"role" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Type      string `json:"type"`
	MediaURL  string `json:"media_url"`
}

// CreateChatSessionRequest 创建聊天会话请求
type CreateChatSessionRequest struct {
	Title string `json:"title" binding:"required"`
}

// ConsumPointsRequest 消费积分请求
type ConsumPointsRequest struct {
	Action string `json:"action" binding:"required"`
}

// SubscribePlanRequest 订阅计划请求
type SubscribePlanRequest struct {
	Plan string `json:"plan" binding:"required"`
}

// BuyCreditsRequest 購买积分请求
type BuyCreditsRequest struct {
	PackageID string `json:"package_id" binding:"required"`
}

// WidgetDefinition 小组件定义
type WidgetDefinition struct {
	ID             string `gorm:"type:varchar(191);primaryKey" json:"id" comment:"小组件ID"`
	Name           string `gorm:"type:varchar(128);not null" json:"name" comment:"小组件名称"`
	Description    string `gorm:"type:varchar(512)" json:"description" comment:"小组件描述"`
	Icon           string `gorm:"type:varchar(32)" json:"icon" comment:"图标"`
	Category       string `gorm:"type:varchar(64)" json:"category" comment:"分类"`
	SupportedSizes string `gorm:"type:varchar(128)" json:"supported_sizes" comment:"支持的尺寸"`
	PreviewColor   string `gorm:"type:varchar(64)" json:"preview_color" comment:"预览颜色"`
	ComponentType  string `gorm:"type:varchar(64)" json:"component_type" comment:"组件类型"`
}

// GetCurrentTime 获取当前时间
func GetCurrentTime() time.Time {
	return time.Now()
}
