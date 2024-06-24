package implements

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/entity"
	"log/slog"
	"net/http"
)

type WebRestService struct {
	Url string
}

func NewWebRestService(url string) *WebRestService {
	return &WebRestService{Url: url}
}

func (ws *WebRestService) Send(ctx context.Context, notifications []entity.Notification) error {
	bodyBytes, err := json.Marshal(notifications)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	reader := bytes.NewReader(bodyBytes)
	resp, err := http.DefaultClient.Post(ws.Url, "application/json", reader)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Error(resp.Status)
		return errors.New(resp.Status)
	}
	return nil
}
