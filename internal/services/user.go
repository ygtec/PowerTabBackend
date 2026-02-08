package services

import (
	"errors"

	"gorm.io/gorm"
	"powertab/config"
	"powertab/internal/models"
	"powertab/internal/utils"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{db: config.GetDB()}
}

func (s *UserService) GetPreferences(userID string) (*models.UserPreference, error) {
	var pref models.UserPreference
	if err := s.db.Where("user_id = ?", userID).First(&pref).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pref = models.UserPreference{
				ID:     utils.GenerateID(),
				UserID: userID,
				Theme:  "auto",
			}
			s.db.Create(&pref)
			return &pref, nil
		}
		return nil, err
	}
	return &pref, nil
}

func (s *UserService) UpdatePreferences(userID string, updates *models.PreferencesUpdateRequest) (*models.UserPreference, error) {
	pref, err := s.GetPreferences(userID)
	if err != nil {
		return nil, err
	}

	if updates.Theme != "" {
		pref.Theme = updates.Theme
	}
	if updates.Language != "" {
		pref.Language = updates.Language
	}
	if updates.DefaultEngine != "" {
		pref.DefaultEngine = updates.DefaultEngine
	}
	if updates.WallpaperURL != "" {
		pref.WallpaperURL = updates.WallpaperURL
	}
	if updates.IconSize != "" {
		pref.IconSize = updates.IconSize
	}
	pref.ShowSeconds = updates.ShowSeconds
	pref.Notifications = updates.Notifications
	pref.AutoWallpaper = updates.AutoWallpaper

	if err := s.db.Save(pref).Error; err != nil {
		return nil, err
	}
	return pref, nil
}

func (s *UserService) ListSearchEngines(userID string) ([]models.SearchEngine, error) {
	var engines []models.SearchEngine
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&engines).Error; err != nil {
		return nil, err
	}
	return engines, nil
}

func (s *UserService) AddSearchEngine(userID string, req *models.SearchEngineRequest) (*models.SearchEngine, error) {
	engine := &models.SearchEngine{
		ID:     utils.GenerateID(),
		UserID: userID,
		Name:   req.Name,
		URL:    req.URL,
	}
	if err := s.db.Create(engine).Error; err != nil {
		return nil, err
	}
	return engine, nil
}

func (s *UserService) DeleteSearchEngine(userID, engineID string) error {
	return s.db.Where("id = ? AND user_id = ?", engineID, userID).Delete(&models.SearchEngine{}).Error
}

func (s *UserService) ListWallpapers(userID string) ([]models.SavedWallpaper, error) {
	var wallpapers []models.SavedWallpaper
	if err := s.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&wallpapers).Error; err != nil {
		return nil, err
	}
	return wallpapers, nil
}

func (s *UserService) AddWallpaper(userID string, req *models.WallpaperRequest) (*models.SavedWallpaper, error) {
	wallpaper := &models.SavedWallpaper{
		ID:     utils.GenerateID(),
		UserID: userID,
		URL:    req.URL,
	}
	if err := s.db.Create(wallpaper).Error; err != nil {
		return nil, err
	}
	return wallpaper, nil
}

func (s *UserService) DeleteWallpaper(userID, wallpaperID string) error {
	return s.db.Where("id = ? AND user_id = ?", wallpaperID, userID).Delete(&models.SavedWallpaper{}).Error
}
