package cache

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

type RedisClient struct {
	Client *redis.Pool
}

var c = &RedisClient{}

func init() {
	c = NewRedisClient("codis url")
}

func Client() *RedisClient {
	return c
}

func NewRedisClient(addr string) *RedisClient {
	rc := new(RedisClient)
	rc.initialize(addr)
	return rc
}

func (rc *RedisClient) initialize(addr string) {
	rc.Client = &redis.Pool{
		MaxIdle:     512,
		MaxActive:   1024,
		IdleTimeout: 0,
		Dial: func() (redis.Conn, error) {
			option := redis.DialPassword("password")
			c, err := redis.Dial("tcp", addr, option)
			if err != nil {
				log.Printf("%s", err.Error())
				return nil, err
			}
			return c, nil
		},
	}
}
