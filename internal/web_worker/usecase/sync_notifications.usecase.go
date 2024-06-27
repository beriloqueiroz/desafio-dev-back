package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/usecase/interfaces"
	"log/slog"
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
