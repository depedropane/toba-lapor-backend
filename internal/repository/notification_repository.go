package repository

import (
	"toba-lapor-backend/internal/model"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification *model.Notification) error
	FindByUserID(userID uint) ([]model.Notification, error)
	FindByID(id uint) (*model.Notification, error)
	Update(notification *model.Notification) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db}
}

func (r *notificationRepository) Create(notification *model.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) FindByUserID(userID uint) ([]model.Notification, error) {
	var notifs []model.Notification
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&notifs).Error
	return notifs, err
}

func (r *notificationRepository) FindByID(id uint) (*model.Notification, error) {
	var notif model.Notification
	err := r.db.First(&notif, id).Error
	if err != nil {
		return nil, err
	}
	return &notif, nil
}

func (r *notificationRepository) Update(notification *model.Notification) error {
	return r.db.Save(notification).Error
}
