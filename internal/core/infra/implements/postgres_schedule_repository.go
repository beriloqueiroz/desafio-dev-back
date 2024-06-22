package implements

import (
	"context"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"github.com/google/uuid"
	"time"
)

type PostgresScheduleRepository struct{}

func (p *PostgresScheduleRepository) Save(ctx context.Context, scheduleNotification *entity.ScheduleNotification) error {
	//TODO implement me
	return nil
}

func (p *PostgresScheduleRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	return nil
}

func (p *PostgresScheduleRepository) FindFirstNotExecutedBeforeDate(ctx context.Context, date time.Time) (*entity.ScheduleNotification, error) {
	//TODO implement me
	return entity.NewScheduleNotification(uuid.NewString(), time.Now().Add(-1*time.Hour), false)
}
