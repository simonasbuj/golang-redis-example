package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-redis-example/internal/cache"
	"log/slog"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)


func main() {

	rs, err := cache.NewRedisCache(&redis.Options{
		Addr: 		"localhost:6379",
		Password: 	"",
		DB: 		0,

	})
	if err != nil {
		slog.Error(err.Error())
	}

	// values to cache
	userID := uuid.NewString()
	jsonUser, _ := json.Marshal(User{
		ID:			userID,
		Username: 	"sim sim",
		Email: 		"sim@one.lt",
		Age:		43,

	})
	userKey := fmt.Sprintf("user:%s", userID)

	err = rs.SetValue(context.Background(), userKey, string(jsonUser), 0)
	if err != nil {return}

	val, err := rs.GetValue(context.Background(), userKey)
	if err != nil {return}

	slog.Info("value retrieved", "value", val)

}

type User struct {
	ID 			string
	Username 	string
	Email		string
	Age			int
}
