package monica

import (
	"fmt"
	"gopkg.in/redis.v3"
	"log"
)

type MonicaRedis struct {
	RedisClient *redis.Client
}

func NewRedisClient(server *MonicaServer) *MonicaRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     server.Config.RedisSchema,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println(server.Config.RedisSchema)
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("ping error")
	}

	fmt.Println(pong)

	return &MonicaRedis{
		RedisClient: client,
	}
}
