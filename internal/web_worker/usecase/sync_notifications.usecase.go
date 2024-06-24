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
	batchSize := 10 // pode ser uma variável de ambiente
	ch := make(chan []entity.Notification, batchSize)
	err := u.NotificationQueues.Read(ctx, ch)
	if err != nil {
		return err
	}
	for notifications := range ch {
		err = u.WebService.Send(ctx, notifications)
		if err != nil {
			slog.Error(err.Error())
			continue
		}
	}
	return nil
}
