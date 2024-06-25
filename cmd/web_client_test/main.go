package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type InputDto struct {
	ID   string `json:"ID"`
	User struct {
		ID         string `json:"ID"`
		Active     bool   `json:"Active"`
		Email      string `json:"Email"`
		Phone      string `json:"Phone"`
		Location   string `json:"City"`
		CreateTime string `json:"CreateTime"`
	} `json:"User"`
	ScheduleNotification struct {
		ID        string    `json:"ID"`
		StartTime time.Time `json:"StartTime"`
		Status    string    `json:"Status"`
	} `json:"ScheduleNotification"`
	Message string `json:"Message"`
}

type output struct {
	Message string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		var input []InputDto
		err := json.NewDecoder(r.Body).Decode(&input)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&output{
				Message: err.Error(),
			})
			return
		}
		clientRedis := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "redis123",
			DB:       0,
		})
		ctx := context.Background()
		fmt.Println(clientRedis.Get(ctx, "São Paulo - SP"))
		fmt.Println(clientRedis.Get(ctx, "Fortaleza - CE"))
		fmt.Println(clientRedis.Get(ctx, "Quixadá - CE"))
		fmt.Println(clientRedis.Get(ctx, "Baturité - CE"))
		//fmt.Println("Input: ", len(input), input)
		json.NewEncoder(w).Encode(input)
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Listening on :9000")
	http.ListenAndServe(":9000", mux)
}
