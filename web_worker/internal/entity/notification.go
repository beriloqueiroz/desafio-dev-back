package entity

import (
	"fmt"
	"time"
)

type Notification struct {
	ID                   string
	User                 User
	ScheduleNotification ScheduleNotification
	Message              string
}

type User struct {
	ID         string
	Active     bool
	Email      string
	Phone      string
	Location   Location
	CreateTime time.Time
}

type ScheduleNotification struct {
	ID        string
	StartTime time.Time
	Status    string
}

type Location struct {
	City  string
	State string
}

func (l Location) String() string {
	return fmt.Sprintf("%s - %s", l.City, l.State)
}
