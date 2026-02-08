package services

import (
	"time"

	"powertab/config"
	"powertab/internal/models"
	"powertab/internal/utils"

	"gorm.io/gorm"
)

type BillingService struct {
	db *gorm.DB
}

func NewBillingService() *BillingService {
	return &BillingService{db: config.GetDB()}
}

func (s *BillingService) ConsumePoints(userID string, req *models.ConsumPointsRequest) (*models.User, error) {
	var user models.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	pointsMap := map[string]int{
		"chat_std":    10,
		"chat_pro":    20,
		"image_gen":   50,
		"video_gen":   100,
		"ppt_gen":     30,
		"mindmap_gen": 20,
		"id_photo":    15,
	}

	points, exists := pointsMap[req.Action]
	if !exists {
		points = 10
	}

	if user.Points < points {
		return nil, gorm.ErrInvalidData
	}

	user.Points -= points
	s.db.Model(&user).Update("points", user.Points)

	record := &models.PointConsumption{
		ID:     utils.GenerateID(),
		UserID: userID,
		Action: req.Action,
		Points: points,
	}
	s.db.Create(record)

	return &user, nil
}

func (s *BillingService) SubscribePlan(userID string, req *models.SubscribePlanRequest) (*models.User, error) {
	var user models.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	user.PlanLevel = req.Plan
	now := time.Now()
	expiry := now.AddDate(0, 1, 0)

	switch req.Plan {
	case "pro":
		user.PlanExpiry = &expiry
		user.Points = 500
	case "ultra":
		user.PlanExpiry = &expiry
		user.Points = 1500
	default:
		user.PlanLevel = "free"
	}

	s.db.Save(&user)

	return &user, nil
}

func (s *BillingService) BuyCredits(userID string, req *models.BuyCreditsRequest) (*models.User, error) {
	var user models.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	creditsMap := map[string]int{
		"package_1":  100,
		"package_5":  600,
		"package_10": 1300,
	}

	credits, exists := creditsMap[req.PackageID]
	if !exists {
		credits = 100
	}

	user.Points += credits
	s.db.Save(&user)

	return &user, nil
}
