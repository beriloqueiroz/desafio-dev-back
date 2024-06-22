package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewScheduleNotification(t *testing.T) {
	schedule, err := NewScheduleNotification(uuid.NewString(), time.Now().Add(time.Hour*24), false)
	assert.Nil(t, err)
	assert.NotNil(t, schedule)
}

func TestNewScheduleNotification_TimeInPast(t *testing.T) {
	schedule, err := NewScheduleNotification(uuid.NewString(), time.Now().Add(-time.Hour*24), false)
	assert.NotNil(t, err)
	assert.Nil(t, schedule)
	assert.Equal(t, "start time is invalid", err.Error())
}

func TestNewScheduleNotification_InvalidID(t *testing.T) {
	schedule, err := NewScheduleNotification("123", time.Now().Add(time.Hour*24), false)
	assert.NotNil(t, err)
	assert.Nil(t, schedule)
	assert.Equal(t, "id is invalid", err.Error())
}

func TestScheduleNotification_Execute(t *testing.T) {
	schedule, err := NewScheduleNotification(uuid.NewString(), time.Now().Add(time.Hour*24), false)
	assert.Nil(t, err)
	schedule.Execute()
	assert.Equal(t, true, schedule.Executed)
}
