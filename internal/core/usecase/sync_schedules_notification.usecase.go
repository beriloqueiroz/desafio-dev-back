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
	scheduler, err := u.ScheduleRepository.FindFirstNotExecutedBeforeDate(ctx, time.Now())
	if scheduler == nil && err == nil {
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Iniciando schedule: %s", scheduler.ID))
	// todo buscar users ativos com paginação
	page := 1
	size := 500
	for {
		users, err := u.UserRepository.ListActives(ctx, page, size)
		if users == nil || err != nil {
			msg := fmt.Sprintf("Sem usuários para schedule,  schedule: %s", scheduler.ID)
			fmt.Println(msg)
			scheduler.MarkStatus(msg)
			break
		}
		if err != nil {
			fmt.Println(err)
			scheduler.MarkStatus(fmt.Sprintf("Falha ao lista users,  page: %d, error: %s", page, err.Error()))
			break
		}
		uniquesLocations := getUniquesLocation(users)
		// todo buscar mensagens com base nas cidades dos usuários
		locationsMapMsg, err := u.MessageRepository.ListByLocations(ctx, uniquesLocations)
		if err != nil {
			fmt.Println(err)
			scheduler.MarkStatus(fmt.Sprintf("Falha ao lista locations,  page: %d, error: %s", page, err.Error()))
			break
		}
		var msgErr string
		for _, user := range users {
			// todo montar notificações enviar notificações para as filas
			notification, err := entity.NewNotification(uuid.NewString(), user, *scheduler, locationsMapMsg[user.Location])
			if err != nil {
				fmt.Println(err)
				msgErr += fmt.Sprintf("Falha ao enviar notificação,  page: %d, id: %s, user: %s, error: %s;", page, notification.ID, user.ID, err.Error())
				scheduler.MarkStatus(msgErr)
				continue
			}
			for _, queue := range u.NotificationQueues {
				err = queue.Send(ctx, notification)
				if err != nil {
					fmt.Println(err)
					msgErr += fmt.Sprintf("Falha ao enviar notificação,  page: %d, id: %s, user: %s, error: %s;", page, notification.ID, user.ID, err.Error())
					scheduler.MarkStatus(msgErr)
				}
			}

		}
		if len(users) < size {
			break
		}
		page++
	}
	// todo marcar schedulers como executed
	scheduler.Execute()
	u.ScheduleRepository.Save(ctx, scheduler)
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
