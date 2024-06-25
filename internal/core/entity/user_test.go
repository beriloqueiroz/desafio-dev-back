package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(uuid.NewString(), true, "test@test.com.br", "11888888888", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	assert.Nil(t, err)
	assert.Equal(t, "test@test.com.br", user.Email)
	assert.Equal(t, "11888888888", user.Phone)
}

func TestUser_Activate(t *testing.T) {
	user, _ := NewUser(uuid.NewString(), false, "test@test.com.br", "11888888888", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	user.Activate()
	assert.Equal(t, true, user.Active)
}

func TestUser_Deactivate(t *testing.T) {
	user, _ := NewUser(uuid.NewString(), true, "test@test.com.br", "11888888888", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	user.Deactivate()
	assert.Equal(t, false, user.Active)
}

func TestUser_IsActive(t *testing.T) {
	user, _ := NewUser(uuid.NewString(), true, "test@test.com.br", "11888888888", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	assert.Equal(t, true, user.Active)
}

func TestUser_InvalidID(t *testing.T) {
	user, err := NewUser("123", true, "test@test.com.br", "11888888888", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "id is invalid")
	assert.Nil(t, user)
}

func TestUser_InvalidEmail(t *testing.T) {
	user, err := NewUser(uuid.NewString(), true, "", "11888888888", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "email is required")
	assert.Nil(t, user)
}

func TestUser_InvalidPhone(t *testing.T) {
	user, err := NewUser(uuid.NewString(), true, "test@test.com.br", "", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "phone is required")
	assert.Nil(t, user)
}

func TestUser_InvalidPhoneAndEmail(t *testing.T) {
	user, err := NewUser(uuid.NewString(), true, "", "", Location{
		City:  "Fortaleza",
		State: "CE",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "phone is required; email is required")
	assert.Nil(t, user)
}

func TestUser_InvalidState(t *testing.T) {
	user, err := NewUser(uuid.NewString(), true, "test@test.com.br", "85989898989", Location{
		City:  "Fortaleza",
		State: "Ceará",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "location state is invalid")
	assert.Nil(t, user)
}

func TestUser_InvalidCity(t *testing.T) {
	user, err := NewUser(uuid.NewString(), true, "test@test.com.br", "85989898989", Location{
		City:  "",
		State: "CE",
	})
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "location city is required")
	assert.Nil(t, user)
}

func TestLocation_String(t *testing.T) {
	l, err := LocationByString("São Paulo - SP")
	assert.Nil(t, err)
	assert.Equal(t, "SP", l.State)
	assert.Equal(t, "São Paulo", l.City)
	l, err = LocationByString("Rio-de-janeiro - RJ")
	assert.Nil(t, err)
	assert.Equal(t, "RJ", l.State)
	assert.Equal(t, "Rio-de-janeiro", l.City)
}
