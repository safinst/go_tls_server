package cache

import (
	"errors"
	"fmt"
)

var ErrRedisNotByte = errors.New("get_result_is_not_byte")

//Close
func (rc *RedisClient) Close() error {
	return rc.Client.Close()
}

// produce key
func MakeKey(prefix string, key interface{}) string {
	return fmt.Sprintf("%s%v", prefix, key)
}

// redis set command
func (rc *RedisClient) Set(key string, value interface{}) error {
	c := rc.Client.Get()
	defer c.Close()
	_, err := c.Do("SET", key, value)
	return err
}

// redis SETEX command
func (rc *RedisClient) SetExpire(key string, time int32, value interface{}) error {
	c := rc.Client.Get()
	defer c.Close()
	_, err := c.Do("SETEX", key, time, value)
	return err
}

// redis get command
func (rc *RedisClient) Get(key string) ([]byte, error) {
	c := rc.Client.Get()
	defer c.Close()
	value, err := c.Do("GET", key)
	if err != nil {
		return nil, err
	}
	if data, ok := value.([]byte); ok {
		return data, nil
	}
	return nil, ErrRedisNotByte
}
