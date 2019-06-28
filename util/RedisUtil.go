package util

import "github.com/go-redis/redis"

func main() {
	// 内部维护了一个连接池
	var client = redis.NewClient(&redis.Options{
		Addr: "47.99.145.78:6377",
	})

	client.Get("test").Result()
}
