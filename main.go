package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1234",
		DB:       0,
	})

	ctx := context.Background()

	soundFiles := []string{"sound1.wav", "sound2.mp3", "sound3.flac"}
	for _, file := range soundFiles {
		err := client.RPush(ctx, "queue", file).Err()
		if err != nil {
			log.Fatalf("Error adding %s to queue: %v\n", file, err)
		}
	}
}
