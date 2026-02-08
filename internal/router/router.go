package router

import (
	"powertab/internal/handlers"
	"powertab/internal/middleware"
	"powertab/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置所有路由
func SetupRouter(router *gin.Engine) {
	// 应用中间件
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandlingMiddleware())

	// 初始化服务
	authSvc := services.NewAuthService()
	userSvc := services.NewUserService()
	layoutSvc := services.NewLayoutService()
	calendarSvc := services.NewCalendarService()
	aiSvc := services.NewAIService()
	billingSvc := services.NewBillingService()
	initDataSvc := services.NewInitDataService()

	// 初始化处理器
	authHandler := handlers.NewAuthHandler(authSvc)
	userHandler := handlers.NewUserHandler(userSvc)
	layoutHandler := handlers.NewLayoutHandler(layoutSvc)
	calendarHandler := handlers.NewCalendarHandler(calendarSvc)
	aiHandler := handlers.NewAIHandler(aiSvc)
	billingHandler := handlers.NewBillingHandler(billingSvc)
	dataInitHandler := handlers.NewDataInitHandler(initDataSvc)

	// API v1 路由组
	api := router.Group("/api")

	// 认证路由 (不需要认证)
	setupAuthRoutes(api, authHandler)

	// 用户路由 (需要认证)
	setupUserRoutes(api, authHandler, userHandler)

	// 布局路由 (需要认证)
	setupLayoutRoutes(api, layoutHandler)

	// 日历路由 (需要认证)
	setupCalendarRoutes(api, calendarHandler)

	// AI 路由 (需要认证)
	setupAIRoutes(api, aiHandler)

	// 计费路由 (需要认证)
	setupBillingRoutes(api, billingHandler)

	// 数据路由 (需要认证)
	setupDataRoutes(api, dataInitHandler)

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "OK",
			"data": nil,
		})
	})
}

// setupAuthRoutes 认证路由配置
func setupAuthRoutes(api *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
	}
}

// setupUserRoutes 用户路由配置
func setupUserRoutes(api *gin.RouterGroup, authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler) {
	user := api.Group("/user", middleware.AuthMiddleware())
	{
		user.GET("/current", authHandler.GetCurrentUser)
		user.POST("/update", authHandler.UpdateUser)
		user.GET("/preferences", userHandler.GetPreferences)
		user.POST("/preferences/update", userHandler.UpdatePreferences)
		user.GET("/search_engines", userHandler.ListSearchEngines)
		user.POST("/search_engines/add", userHandler.AddSearchEngine)
		user.POST("/search_engines/delete", userHandler.DeleteSearchEngine)
		user.GET("/wallpapers", userHandler.ListWallpapers)
		user.POST("/wallpapers/add", userHandler.AddWallpaper)
		user.POST("/wallpapers/delete", userHandler.DeleteWallpaper)
	}
}

// setupLayoutRoutes 布局路由配置
func setupLayoutRoutes(api *gin.RouterGroup, layoutHandler *handlers.LayoutHandler) {
	layout := api.Group("/layout", middleware.AuthMiddleware())
	{
		layout.GET("/categories", layoutHandler.GetCategories)
		layout.POST("/categories/save", layoutHandler.SaveCategories)
		layout.GET("/speed_dial", layoutHandler.GetSpeedDial)
		layout.POST("/speed_dial/save", layoutHandler.SaveSpeedDial)
		layout.GET("/widgets", layoutHandler.GetWidgets)
		layout.POST("/widgets/save", layoutHandler.SaveWidgets)
		layout.GET("/recommended_tools", layoutHandler.GetRecommendedTools)
		layout.POST("/recommended_tools/save", layoutHandler.SaveRecommendedTools)
		layout.GET("/tool_order", layoutHandler.GetToolOrder)
		layout.POST("/tool_order/save", layoutHandler.SaveToolOrder)
	}
}

// setupCalendarRoutes 日历路由配置
func setupCalendarRoutes(api *gin.RouterGroup, calendarHandler *handlers.CalendarHandler) {
	calendar := api.Group("/calendar", middleware.AuthMiddleware())
	{
		calendar.GET("/events", calendarHandler.GetEvents)
		calendar.POST("/events/add", calendarHandler.AddEvent)
		calendar.POST("/events/update", calendarHandler.UpdateEvent)
		calendar.POST("/events/delete", calendarHandler.DeleteEvent)
	}
}

// setupAIRoutes AI 路由配置
func setupAIRoutes(api *gin.RouterGroup, aiHandler *handlers.AIHandler) {
	ai := api.Group("/ai", middleware.AuthMiddleware())
	{
		ai.GET("/sessions", aiHandler.GetSessions)
		ai.POST("/sessions/save", aiHandler.SaveSessions)
	}
}

// setupBillingRoutes 计费路由配置
func setupBillingRoutes(api *gin.RouterGroup, billingHandler *handlers.BillingHandler) {
	billing := api.Group("/billing", middleware.AuthMiddleware())
	{
		billing.POST("/consume", billingHandler.ConsumePoints)
		billing.POST("/subscribe", billingHandler.Subscribe)
		billing.POST("/buy_credits", billingHandler.BuyCredits)
	}
}

// setupDataRoutes 数据路由配置
func setupDataRoutes(api *gin.RouterGroup, dataInitHandler *handlers.DataInitHandler) {
	data := api.Group("/data", middleware.AuthMiddleware())
	{
		data.POST("/usage", func(c *gin.Context) {
			c.JSON(200, gin.H{"code": 0, "msg": "ok"})
		})
		data.POST("/initialize", dataInitHandler.InitializeUserData)
		data.GET("/widgets/definitions", dataInitHandler.GetAllWidgetDefinitions)
		data.GET("/categories/default", dataInitHandler.GetDefaultCategories)
		data.GET("/speed-dial/default", dataInitHandler.GetDefaultSpeedDial)
	}
}
