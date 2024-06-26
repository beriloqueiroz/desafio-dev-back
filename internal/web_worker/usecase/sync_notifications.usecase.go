package usecase

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/entity"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/usecase/interfaces"
	"log/slog"
)

type SyncNotificationUseCase struct {
	WebService         interfaces.WebService
	NotificationQueues interfaces.NotificationQueueRepository
}

func (u *SyncNotificationUseCase) Execute(ctx context.Context) error {
	// ler fila
	ch := make(chan []entity.Notification, 1)
	go func() {
		err := u.NotificationQueues.Read(ctx, ch)
		if err != nil {
			slog.Error(err.Error())
		}
	}()
	for notifications := range ch {
		err := u.WebService.Send(ctx, notifications)
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		err = u.NotificationQueues.Commit(ctx)
		if err != nil {
			slog.Error(err.Error())
		}
	}
	return nil
}
