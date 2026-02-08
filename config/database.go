package config

import (
	"fmt"
	"log"

	"powertab/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate 自动迁移所有模型
	err = DB.AutoMigrate(
		&models.User{},
		&models.UserPreference{},
		&models.SearchEngine{},
		&models.SavedWallpaper{},
		&models.Category{},
		&models.SpeedDialItem{},
		&models.Widget{},
		&models.RecommendedTool{},
		&models.ToolOrder{},
		&models.CalendarEvent{},
		&models.ChatMessage{},
		&models.ChatSession{},
		&models.PointConsumption{},
		&models.ToolUsageLog{},
	)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Database initialized successfully")
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}

// CreateDSN 创建 MySQL DSN
func CreateDSN(username, password, host, port, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)
}
