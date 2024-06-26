package implements

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/beriloqueiroz/desafio-dev-back/internal/web_worker/entity"
	"github.com/sony/gobreaker/v2"
	"log/slog"
	"net/http"
)

type WebRestService struct {
	Url string
}

func NewWebRestService(url string) *WebRestService {
	initCircuitBreak()
	return &WebRestService{Url: url}
}

func (ws *WebRestService) Send(ctx context.Context, notifications []entity.Notification) error {
	bodyBytes, err := json.Marshal(notifications)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	reader := bytes.NewReader(bodyBytes)
	return sendWithCircuitBreaker(ws.Url, reader)
}

func send(url string, reader *bytes.Reader) error {
	resp, err := http.DefaultClient.Post(url, "application/json", reader)
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

func initCircuitBreak() {
	var st gobreaker.Settings
	st.Name = "SEND POST WEB APP"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 5 && failureRatio >= 0.5 // todo pode ser variável de ambiente
	}

	cb = gobreaker.NewCircuitBreaker[[]byte](st)
}

var cb *gobreaker.CircuitBreaker[[]byte]

func sendWithCircuitBreaker(url string, reader *bytes.Reader) error {
	_, err := cb.Execute(func() ([]byte, error) {
		return nil, send(url, reader)
	})
	if err != nil {
		if cb.State().String() == "open" {
			return errors.Join(err, errors.New("Circuit breaker is open"))
		}
		return err
	}
	return nil
}
