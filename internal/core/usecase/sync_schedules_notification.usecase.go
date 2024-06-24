package usecase

import (
	"context"
	"fmt"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase/interfaces"
	"github.com/google/uuid"
	"time"
)

type SyncSchedulesNotificationUseCase struct {
	UserRepository     interfaces.UserRepository
	ScheduleRepository interfaces.ScheduleNotificationRepository
	NotificationQueues []interfaces.NotificationQueueRepository
	MessageRepository  interfaces.MessageRepository
}

func (u *SyncSchedulesNotificationUseCase) Execute(ctx context.Context) error {
	// todo buscar primeiro scheduler não executados com data anterior a atual
	scheduler, err := u.ScheduleRepository.FindFirstPendingBeforeDate(ctx, time.Now())
	if scheduler == nil && err == nil {
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Iniciando schedule: %s", scheduler.ID))
	scheduler.MarkProcessing()
	err = u.ScheduleRepository.Save(ctx, scheduler)
	if err != nil {
		return err
	}
	// todo buscar users ativos com paginação
	page := 1
	size := 500
	for {
		users, err := u.UserRepository.ListActives(ctx, page, size)
		if users == nil && err == nil {
			fmt.Println(fmt.Sprintf("Sem usuários para schedule,  schedule: %s", scheduler.ID))
			scheduler.MarkExecuted()
			break
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("Falha ao lista users,  page: %d, error: %s", page, err.Error()))
			scheduler.MarkExecutedWithError()
			break
		}
		uniquesLocations := getUniquesLocation(users)
		// todo buscar mensagens com base nas cidades dos usuários
		locationsMapMsg, err := u.MessageRepository.ListByLocations(ctx, uniquesLocations)
		if err != nil {
			fmt.Println(fmt.Sprintf("Falha ao lista locations,  page: %d, error: %s", page, err.Error()))
			scheduler.MarkExecutedWithError()
			break
		}
		for _, user := range users {
			// todo montar notificações enviar notificações para as filas
			notification, err := entity.NewNotification(uuid.NewString(), user, *scheduler, locationsMapMsg[user.Location])
			if err != nil {
				fmt.Println(fmt.Sprintf("Falha ao enviar notificação,  page: %d, id: %s, user: %s, error: %s;", page, notification.ID, user.ID, err.Error()))
				scheduler.MarkExecutedWithError()
				continue
			}
			for _, queue := range u.NotificationQueues {
				err = queue.Send(ctx, notification)
				if err != nil {
					fmt.Println(fmt.Sprintf("Falha ao enviar notificação,  page: %d, id: %s, user: %s, error: %s;", page, notification.ID, user.ID, err.Error()))
					scheduler.MarkExecutedWithError()
				}
			}

		}
		if len(users) < size {
			if scheduler.Status == entity.Pending {
				scheduler.MarkExecuted()
			}
			break
		}
		page++
	}
	// todo marcar schedulers como executed
	err = u.ScheduleRepository.Save(ctx, scheduler)
	if err != nil {
		return err
	}
	return nil
}

func getUniquesLocation(sliceList []entity.User) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range sliceList {
		if _, value := allKeys[item.Location]; !value {
			allKeys[item.Location] = true
			list = append(list, item.Location)
		}
	}
	return list
}
