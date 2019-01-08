package main

import (
	"os"

	"github.com/go-redis/redis"
)

const (
	DEFAULT_REDIS_ADDRESS = "localhost:6379"
)

type RedisDataStore struct {
	client *redis.Client
}

func (self *RedisDataStore) Put(id, value string) error {
	err := self.client.Set(id, value, 0).Err()
	return err
}

func (self *RedisDataStore) Get(id string) (string, error) {
	value, err := self.client.Get(id).Result()
	return value, err
}

func GetRedisAddress() string {
	addr := os.Getenv("REDIS_URL")
	if addr == "" {
		return addr
	} else {
		return DEFAULT_REDIS_ADDRESS
	}
}

func NewRedisDataStore(address string) (DataStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}

	return &RedisDataStore{client: client}, nil
}
