package entity

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
	"time"
)

type Location struct {
	City  string
	State string
}

func (l Location) String() string {
	return fmt.Sprintf("%s - %s", l.City, l.State)
}

func LocationByString(s string) (Location, error) {
	pattern := `\s-\s[A-Z]{2}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(s) {
		return Location{}, errors.New("invalid location, only format: 'city name - UF'")
	}
	state := s[len(s)-2:]
	return Location{
		State: state,
		City:  strings.ReplaceAll(s, fmt.Sprintf(" - %s", state), ""),
	}, nil
}

type User struct {
	ID         string
	Active     bool
	Email      string
	Phone      string
	Location   Location
	CreateTime time.Time
}

func NewUser(id string, active bool, email string, phone string, location Location) (*User, error) {
	user := &User{
		ID:         id,
		Phone:      phone,
		Email:      email,
		CreateTime: time.Now(),
		Active:     active,
		Location:   location,
	}
	if user.Validate() != nil {
		return nil, user.Validate()
	}
	return user, nil
}

func (u *User) Activate() {
	u.Active = true
}

func (u *User) Deactivate() {
	u.Active = false
}

func (u *User) Validate() error {
	var msg []string
	if uuid.Validate(u.ID) != nil {
		msg = append(msg, "id is invalid")
	}
	if u.Phone == "" {
		msg = append(msg, "phone is required")
	}
	if u.Email == "" {
		msg = append(msg, "email is required")
	}
	if len(u.Location.State) != 2 {
		msg = append(msg, "location state is invalid")
	}
	if u.Location.City == "" {
		msg = append(msg, "location city is required")
	}
	if len(msg) > 0 {
		return errors.New(strings.Join(msg, "; "))
	}
	return nil
}
