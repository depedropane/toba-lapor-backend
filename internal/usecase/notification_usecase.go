package usecase

import (
	"errors"
	"toba-lapor-backend/internal/model/dto"
	"toba-lapor-backend/internal/repository"
)

type NotificationUsecase interface {
	GetMyNotifications(userID uint) ([]dto.NotificationResponse, error)
	MarkAsRead(id uint, userID uint) error
}

type notificationUsecase struct {
	notifRepo repository.NotificationRepository
}

func NewNotificationUsecase(notifRepo repository.NotificationRepository) NotificationUsecase {
	return &notificationUsecase{notifRepo}
}

func (u *notificationUsecase) GetMyNotifications(userID uint) ([]dto.NotificationResponse, error) {
	notifs, err := u.notifRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var res []dto.NotificationResponse
	for _, n := range notifs {
		res = append(res, dto.NotificationResponse{
			ID:        n.ID,
			Title:     n.Title,
			Body:      n.Body,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt,
		})
	}
	return res, nil
}

func (u *notificationUsecase) MarkAsRead(id uint, userID uint) error {
	notif, err := u.notifRepo.FindByID(id)
	if err != nil {
		return errors.New("notification not found")
	}

	if notif.UserID != userID {
		return errors.New("unauthorized to update this notification")
	}

	notif.IsRead = true
	return u.notifRepo.Update(notif)
}
