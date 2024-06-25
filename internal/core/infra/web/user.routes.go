package web

import (
	"encoding/json"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase"
	"log/slog"
	"net/http"
)

type UserRoutes struct {
	InsertUseCase     usecase.InsertUserUseCase
	ActivateUseCase   usecase.ActivateUserUseCase
	DeactivateUseCase usecase.DeactivateUserUseCase
}

func NewUserRoutes(
	insertUseCase usecase.InsertUserUseCase,
	activeUseCase usecase.ActivateUserUseCase,
	useCase usecase.DeactivateUserUseCase) *UserRoutes {
	return &UserRoutes{
		InsertUseCase:     insertUseCase,
		ActivateUseCase:   activeUseCase,
		DeactivateUseCase: useCase,
	}
}

type InsertUserInputDto struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
	City  string `json:"city"`
	State string `json:"state"`
}

type InsertUserOutputDto struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func (cr *UserRoutes) InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var input InsertUserInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&output{
			Message: err.Error(),
		})
		return
	}
	id, err := cr.InsertUseCase.Execute(r.Context(), input.Email, input.Phone, input.City, input.State)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&output{
			Message: err.Error(),
		})
		return
	}
	output := &InsertUserOutputDto{
		Id:      id,
		Message: "User successfully inserted",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (cr *UserRoutes) ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := cr.ActivateUseCase.Execute(r.Context(), id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&output{
			Message: err.Error(),
		})
		return
	}
	output := &output{
		Message: "Activate Success",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (cr *UserRoutes) DeactivateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := cr.DeactivateUseCase.Execute(r.Context(), id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&output{
			Message: err.Error(),
		})
		return
	}
	output := &output{
		Message: "Deactivate Success",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
