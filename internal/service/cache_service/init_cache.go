package cache_service

import (
	"encoding/json"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			c, err := redis.Dial("tcp", global.Setting.Redis.Host)
			if err != nil {
				logging.Log().Fatalf("cache_service.Setup | redis [%s] 连接失败", global.Setting.Redis.Host)
				return nil, err
			}
			if global.Setting.Redis.Password != "" {
				_, err := c.Do("AUTH", global.Setting.Redis.Password)
				if err != nil {
					c.Close()
					logging.Log().Fatal(" cache_service.Setup | redis 连接失败, 密码错误")
					return nil, err
				}
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     global.Setting.Redis.MaxIdle,
		MaxActive:   global.Setting.Redis.MaxActive,
		IdleTimeout: global.Setting.Redis.IdleTimeout,
	}

	return nil
}

func Set(key string, data interface{}) error {
	return SetWithExpire(key, data, 0)
}

// 设置键值对
func SetWithExpire(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)

	}
	return err
}

// 根据 key 获取 value
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	value, err := redis.Bytes(conn.Do("GET", key))
	return value, err
}

// 判断 key 是否存在
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, _ := redis.Bool(conn.Do("Exists", key))
	return exists
}

// 删除某个key
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DELETE", key))
}

// 批量删除包含 key 的键值对
func DeleteLikes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err := Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
