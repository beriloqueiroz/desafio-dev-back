package entity

import (
	"errors"
	"github.com/google/uuid"
	"strings"
	"time"
)

type ScheduleNotification struct {
	ID        string
	StartTime time.Time
	Executed  bool
}

func NewScheduleNotification(id string, startTime time.Time, executed bool) (*ScheduleNotification, error) {
	schedule := &ScheduleNotification{
		ID:        id,
		StartTime: startTime,
		Executed:  executed,
	}
	err := schedule.Validate()
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *ScheduleNotification) Execute() {
	s.Executed = true
}

func (s *ScheduleNotification) Validate() error {
	var msg []string
	if uuid.Validate(s.ID) != nil {
		msg = append(msg, "id is invalid")
	}
	if s.StartTime.Before(time.Now()) {
		msg = append(msg, "start time is invalid")
	}
	if len(msg) > 0 {
		return errors.New(strings.Join(msg, "; "))
	}
	return nil
}
