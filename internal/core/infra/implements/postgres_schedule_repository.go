package implements

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/entity"
	"time"
)

type PostgresScheduleRepository struct {
	Db *sql.DB
}

type Schedule struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"start_time"`
	Executed  bool      `json:"executed"`
	Status    string    `json:"status"`
}

func (p *PostgresScheduleRepository) Save(ctx context.Context, scheduleNotification *entity.ScheduleNotification) error {
	var s Schedule
	fmt.Println(scheduleNotification)
	err := p.Db.QueryRowContext(ctx, "SELECT id, start_time, executed, status FROM schedules WHERE id = $1", scheduleNotification.ID).Scan(
		&s.ID, &s.StartTime, &s.Executed, &s.Status)
	if errors.Is(err, sql.ErrNoRows) {
		// insert
		_, err = p.Db.ExecContext(ctx, "INSERT INTO schedules (id, start_time, executed, status) VALUES ($1, $2, $3, $4)",
			scheduleNotification.ID, scheduleNotification.StartTime, scheduleNotification.Executed, scheduleNotification.Status)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	// update
	_, err = p.Db.ExecContext(ctx, "UPDATE schedules SET start_time = $1, executed = $2, status = $3  WHERE id = $4;",
		scheduleNotification.StartTime, scheduleNotification.Executed, scheduleNotification.Status, scheduleNotification.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresScheduleRepository) Delete(ctx context.Context, id string) error {
	_, err := p.Db.ExecContext(ctx, "DELETE FROM schedules WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresScheduleRepository) FindFirstNotExecutedBeforeDate(ctx context.Context, date time.Time) (*entity.ScheduleNotification, error) {
	var s Schedule
	err := p.Db.QueryRowContext(ctx, "SELECT id, start_time, executed, status FROM schedules WHERE start_time <= $1 and executed=false", date.Format("2006-01-02T15:04:05")).Scan(
		&s.ID, &s.StartTime, &s.Executed, &s.Status)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return entity.NewScheduleNotification(s.ID, s.StartTime, s.Executed)
}
