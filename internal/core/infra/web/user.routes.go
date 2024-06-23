package web

import (
	"encoding/json"
	"net/http"
)

type UserRoutes struct {
}

func NewUserRoutes() *UserRoutes {
	return &UserRoutes{}
}

type output struct {
	Message string
}

func (cr *UserRoutes) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	output := &output{
		Message: "Success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (cr *UserRoutes) ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	output := &output{
		Message: "Activate",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
