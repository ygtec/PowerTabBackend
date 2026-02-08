package services

import (
	"gorm.io/gorm"
	"powertab/config"
	"powertab/internal/models"
	"powertab/internal/utils"
)

type AIService struct {
	db *gorm.DB
}

func NewAIService() *AIService {
	return &AIService{db: config.GetDB()}
}

func (s *AIService) GetSessions(userID string) ([]models.ChatSession, error) {
	var sessions []models.ChatSession
	if err := s.db.Where("user_id = ?", userID).Preload("Messages").Order("updated_at DESC").Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *AIService) CreateSession(userID string, title string) (*models.ChatSession, error) {
	session := &models.ChatSession{
		ID:     utils.GenerateID(),
		UserID: userID,
		Title:  title,
	}
	if err := s.db.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AIService) AddMessage(userID string, sessionID string, req *models.AddChatMessageRequest) (*models.ChatMessage, error) {
	var session models.ChatSession
	if err := s.db.Where("id = ? AND user_id = ?", sessionID, userID).First(&session).Error; err != nil {
		return nil, err
	}

	message := &models.ChatMessage{
		ID:        utils.GenerateID(),
		SessionID: sessionID,
		Role:      req.Role,
		Content:   req.Content,
		Type:      req.Type,
		MediaURL:  req.MediaURL,
	}
	if err := s.db.Create(message).Error; err != nil {
		return nil, err
	}

	s.db.Model(&session).Update("updated_at", "NOW()")
	return message, nil
}

func (s *AIService) SaveSessions(userID string, sessions []models.ChatSession) error {
	for _, session := range sessions {
		session.UserID = userID
		s.db.Save(&session)
	}
	return nil
}
