package web

import (
	"encoding/json"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/usecase"
	"log/slog"
	"net/http"
)

type UserRoutes struct {
	ActivateUseCase   usecase.ActivateUserUseCase
	DeactivateUseCase usecase.DeactivateUserUseCase
}

func NewUserRoutes(
	activeUseCase usecase.ActivateUserUseCase,
	useCase usecase.DeactivateUserUseCase) *UserRoutes {
	return &UserRoutes{
		ActivateUseCase:   activeUseCase,
		DeactivateUseCase: useCase,
	}
}

type InsertUserOutputDto struct {
	Id      string `json:"id"`
	Message string `json:"message"`
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
