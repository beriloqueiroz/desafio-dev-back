package usecase

import (
	"context"
	"log/slog"
	"web_worker/internal/usecase/interfaces"
)

type SyncNotificationUseCase struct {
	WebService         interfaces.WebService
	NotificationQueues interfaces.NotificationQueueRepository
}

func (u *SyncNotificationUseCase) Execute(ctx context.Context) error {
	err := u.NotificationQueues.Read(ctx, u.WebService.Send)
	if err != nil {
		slog.Error(err.Error())
	}
	return nil
}
