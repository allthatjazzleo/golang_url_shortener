package redis

import (
	"os"

	"github.com/go-redis/redis"
)

// RedisEndpoint is Redis endpoint
var RedisEndpoint = redisEndpoint()

// Client is redis client
var Client = redis.NewClient(&redis.Options{
	Addr:     RedisEndpoint + ":6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

// Nil export redis.nil
var Nil = redis.Nil

func redisEndpoint() (r string) {
	if r = os.Getenv("REDIS_ENDPOINT"); r == "" {
		return "localhost"
	} else {
		return
	}
}
