package services

import (
	"errors"
	"time"

	"powertab/config"
	"powertab/internal/models"
	"powertab/internal/utils"

	"gorm.io/gorm"
)

// AuthService 认证服务
type AuthService struct {
	db *gorm.DB
}

// NewAuthService 创建认证服务
func NewAuthService() *AuthService {
	return &AuthService{db: config.GetDB()}
}

// Register 用户注册
func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	// 如果未提供用户名，从邮箱自动生成
	if username == "" {
		// 从邮箱中提取 @ 前的部分
		for i, r := range email {
			if r == '@' {
				username = email[:i]
				break
			}
		}
	}

	// 检查邮箱是否已存在
	var existingUser models.User
	if err := s.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already registered")
	}

	// 检查用户名是否已存在
	if err := s.db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// 生成令牌
	token, err := utils.GenerateToken(utils.GenerateUserID(), 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &models.User{
		ID:       utils.GenerateUserID(),
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Token:    token,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	// 创建用户偏好设置
	userPref := &models.UserPreference{
		ID:     utils.GenerateID(),
		UserID: user.ID,
		Theme:  "auto",
	}
	s.db.Create(userPref)

	// 清除密码字段
	user.Password = ""
	return user, nil
}

// Login 用户登录
func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email not found")
		}
		return nil, err
	}

	// 验证密码
	if !utils.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid password")
	}

	// 生成新令牌
	token, err := utils.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// 更新令牌
	user.Token = token
	s.db.Model(&user).Update("token", token)

	// 清除密码字段
	user.Password = ""
	return &user, nil
}

// GetUserByID 根据 ID 获取用户
func (s *AuthService) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}

// UpdateUser 更新用户信息
func (s *AuthService) UpdateUser(userID string, updates *models.UpdateUserRequest) (*models.User, error) {
	var user models.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	if updates.Username != "" {
		user.Username = updates.Username
	}
	if updates.Avatar != "" {
		user.Avatar = updates.Avatar
	}

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	user.Password = ""
	return &user, nil
}

// Logout 用户登出
func (s *AuthService) Logout(userID string) error {
	return s.db.Model(&models.User{}).Where("id = ?", userID).Update("token", "").Error
}
