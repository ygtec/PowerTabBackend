package services

import (
	"gorm.io/gorm"
	"powertab/config"
	"powertab/internal/models"
	"powertab/internal/utils"
)

type CalendarService struct {
	db *gorm.DB
}

func NewCalendarService() *CalendarService {
	return &CalendarService{db: config.GetDB()}
}

func (s *CalendarService) GetEvents(userID string) ([]models.CalendarEvent, error) {
	var events []models.CalendarEvent
	if err := s.db.Where("user_id = ?", userID).Order("date DESC, time DESC").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (s *CalendarService) AddEvent(userID string, req *models.CalendarEventRequest) (*models.CalendarEvent, error) {
	event := &models.CalendarEvent{
		ID:        utils.GenerateID(),
		UserID:    userID,
		Title:     req.Title,
		Date:      req.Date,
		Time:      req.Time,
		Type:      req.Type,
		Completed: req.Completed,
		Recurring: req.Recurring,
	}
	if err := s.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (s *CalendarService) UpdateEvent(userID string, eventID string, req *models.CalendarEventRequest) (*models.CalendarEvent, error) {
	var event models.CalendarEvent
	if err := s.db.Where("id = ? AND user_id = ?", eventID, userID).First(&event).Error; err != nil {
		return nil, err
	}

	event.Title = req.Title
	event.Date = req.Date
	event.Time = req.Time
	event.Type = req.Type
	event.Completed = req.Completed
	event.Recurring = req.Recurring

	if err := s.db.Save(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (s *CalendarService) DeleteEvent(userID, eventID string) error {
	return s.db.Where("id = ? AND user_id = ?", eventID, userID).Delete(&models.CalendarEvent{}).Error
}
