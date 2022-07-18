package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	redisClient := newRedisClient()
	result, err := redisPing(redisClient)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func newRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return redisClient
}

func redisPing(client *redis.Client) (string, error) {
	result, err := client.Ping().Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
