package services

import (
	"encoding/json"

	"powertab/config"
	"powertab/internal/models"
	"powertab/internal/utils"

	"gorm.io/gorm"
)

type LayoutService struct {
	db *gorm.DB
}

func NewLayoutService() *LayoutService {
	return &LayoutService{db: config.GetDB()}
}

func (s *LayoutService) GetCategories(userID string) ([]models.Category, error) {
	var categories []models.Category
	if err := s.db.Where("user_id = ?", userID).Order("`order` ASC, created_at ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *LayoutService) SaveCategories(userID string, categories []models.Category) error {
	s.db.Where("user_id = ?", userID).Delete(&models.Category{})
	for i, cat := range categories {
		cat.UserID = userID
		cat.Order = i
		s.db.Create(&cat)
	}
	return nil
}

func (s *LayoutService) GetSpeedDial(userID string) ([]models.SpeedDialItem, error) {
	var items []models.SpeedDialItem
	if err := s.db.Where("user_id = ?", userID).Order("`order` ASC, created_at ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (s *LayoutService) SaveSpeedDial(userID string, items []models.SpeedDialItem) error {
	s.db.Where("user_id = ?", userID).Delete(&models.SpeedDialItem{})
	for i, item := range items {
		item.UserID = userID
		item.Order = i
		s.db.Create(&item)
	}
	return nil
}

func (s *LayoutService) GetWidgets(userID string) ([]models.Widget, error) {
	var widgets []models.Widget
	if err := s.db.Where("user_id = ?", userID).Order("`order` ASC, created_at ASC").Find(&widgets).Error; err != nil {
		return nil, err
	}
	return widgets, nil
}

func (s *LayoutService) SaveWidgets(userID string, widgets []models.Widget) error {
	s.db.Where("user_id = ?", userID).Delete(&models.Widget{})
	for i, widget := range widgets {
		widget.UserID = userID
		widget.Order = i
		s.db.Create(&widget)
	}
	return nil
}

func (s *LayoutService) GetRecommendedTools(userID string) ([]models.RecommendedTool, error) {
	var tools []models.RecommendedTool
	if err := s.db.Where("user_id = ?", userID).Order("`order` ASC, created_at ASC").Find(&tools).Error; err != nil {
		return nil, err
	}
	return tools, nil
}

func (s *LayoutService) SaveRecommendedTools(userID string, tools []models.RecommendedTool) error {
	s.db.Where("user_id = ?", userID).Delete(&models.RecommendedTool{})
	for i, tool := range tools {
		tool.UserID = userID
		tool.Order = i
		s.db.Create(&tool)
	}
	return nil
}

func (s *LayoutService) GetToolOrder(userID string) (map[string][]int, error) {
	var orders []models.ToolOrder
	if err := s.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	result := make(map[string][]int)
	for _, order := range orders {
		var toolIds []int
		if err := json.Unmarshal(order.Orders, &toolIds); err == nil {
			result[order.Category] = toolIds
		}
	}
	return result, nil
}

func (s *LayoutService) SaveToolOrder(userID, category string, orders []int) error {
	toolOrder := &models.ToolOrder{
		ID:       utils.GenerateID(),
		UserID:   userID,
		Category: category,
	}

	orderJSON, err := json.Marshal(orders)
	if err != nil {
		return err
	}
	toolOrder.Orders = orderJSON

	var existing models.ToolOrder
	if err := s.db.Where("user_id = ? AND category = ?", userID, category).First(&existing).Error; err == nil {
		return s.db.Model(&existing).Update("orders", toolOrder.Orders).Error
	}

	return s.db.Create(toolOrder).Error
}
