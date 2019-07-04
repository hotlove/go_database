package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var client *redis.Client

var lock sync.Once

func getRedisClient() *redis.Client {

	lock.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "47.99.145.78:6377",
			Password: "accomplish8512",
		})
	})

	return client
}

func RedisGetValue(key string) string {
	redisClient := getRedisClient()

	defer redisClient.Close()

	value, err := redisClient.Get(key).Result()

	if err != nil {
		fmt.Println(err)
	} else if err != redis.Nil {
		fmt.Println(err)
	}

	return value
}

func ReidsSetMap() {
	redisClient := getRedisClient()
}
