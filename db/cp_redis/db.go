package cp_redis

import (
	"app-container-platform/config"
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisDbForWrite *redis.Client
var redisDbForRead *redis.Client

func InitRedisWriteConnection() error {
	redisDbForWrite = redis.NewClient(&redis.Options{
		Addr:     config.RedisServerForWrite,
		Password: config.RedisServerPassword, // no password set
		DB:       0,                          // use default DB
	})

	log.Println("[INFO] Redis Write connection initialized")
	return nil
}

func InitRedisReadConnection() error {
	redisDbForRead = redis.NewClient(&redis.Options{
		Addr:     config.RedisServerForRead,
		Password: config.RedisServerPassword, // no password set
		DB:       0,                          // use default DB
	})

	log.Println("[INFO] Redis Read connection initialized")
	return nil
}

func Set(key string, val string) error {
	return redisDbForWrite.Set(context.Background(), key, val, 0).Err()
}

func Get(key string) (string, error) {
	return redisDbForRead.Get(context.Background(), key).Result()
}

func GetAllKeys() ([]string, error) {
	return redisDbForRead.Keys(context.Background(), "").Result()
}

func Delete(key string) error {
	return redisDbForWrite.Del(context.Background(), key).Err()
}