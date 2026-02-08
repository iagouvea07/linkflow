package handlers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var	ctx = context.Background()

func InitRedis() (*redis.Client, error){
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Password: "",
		DB: 0,
		Protocol: 2,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}

	fmt.Println("Connected to Redis")
	return rdb, nil
}