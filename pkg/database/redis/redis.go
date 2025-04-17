package redis

import "github.com/redis/go-redis/v9"

func RedisConnection(address string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}
