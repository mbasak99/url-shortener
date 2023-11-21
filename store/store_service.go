package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

// Init store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:3001",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Errorf("Error init Redis: %v\n", err)
	}

	fmt.Printf("Redis started successfully: pong message = {%s}\n", pong)

	storeService.redisClient = redisClient
	return storeService
}

func SaveURLMapping(shortURL string, originalURL string, userID string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		fmt.Errorf("Failed to save key URL | Error: %v - shortURL: %s - originalURL: %s userID: - %v\n", err, shortURL, originalURL, userID)
	}
}

func RetrieveInitialURL(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		fmt.Errorf("Failed RetrieveInitialURL url | Error: %v - shortURL: %s\n", err, shortURL)
	}

	return result
}
