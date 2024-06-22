package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewScheduleNotification(t *testing.T) {
	schedule, err := NewScheduleNotification(uuid.NewString(), "previsão do tempo", time.Now().Add(time.Hour*24))
	assert.Nil(t, err)
	assert.Equal(t, "previsão do tempo", schedule.Message)
}

func TestNewScheduleNotification_TimeInPast(t *testing.T) {
	schedule, err := NewScheduleNotification(uuid.NewString(), "previsão do tempo", time.Now().Add(-time.Hour*24))
	assert.NotNil(t, err)
	assert.Nil(t, schedule)
	assert.Equal(t, "start time is invalid", err.Error())
}

func TestNewScheduleNotification_InvalidID(t *testing.T) {
	schedule, err := NewScheduleNotification("123", "previsão do tempo", time.Now().Add(time.Hour*24))
	assert.NotNil(t, err)
	assert.Nil(t, schedule)
	assert.Equal(t, "id is invalid", err.Error())
}
