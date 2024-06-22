package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
	"time"
)

type SyncSchedulesNotificationUseCase struct {
	UserRepository     interfaces.UserRepository
	ScheduleRepository interfaces.ScheduleNotificationRepository
	NotificationQueue  interfaces.NotificationQueueRepository
	MessageRepository  interfaces.MessageRepository
}

func (u *SyncSchedulesNotificationUseCase) Execute(ctx context.Context, Message string, StartTime time.Time) error {
	return nil
}
