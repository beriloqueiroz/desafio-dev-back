package web

import (
	"encoding/json"
	"github.com/beriloqueiroz/desafio-dev-back/internal/core/usecase"
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
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Location string `json:"location"`
}

func (cr *UserRoutes) InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var input InsertUserInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = cr.InsertUseCase.Execute(r.Context(), input.Email, input.Phone, input.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output := &output{
		Message: "Insert Success",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (cr *UserRoutes) ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := cr.ActivateUseCase.Execute(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output := &output{
		Message: "Activate Success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (cr *UserRoutes) DeactivateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := cr.DeactivateUseCase.Execute(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	output := &output{
		Message: "Deactivate Success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
