package components

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var R *redis.Pool

type RedisConfig struct {
	Addr            string
	Password        string
	Db              string
	MaxIdle         int
	MaxActive       int
	IdleTimeout     int
	Wait            bool
	MaxConnLifetime int
}

func NewRedis(cfg *RedisConfig) error {
	R = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Addr)
			if err != nil {
				return nil, err
			}

			if cfg.Password != "" {
				if _, err := c.Do("AUTH", cfg.Password); err != nil {
					c.Close()
					return nil, err
				}
			}

			if _, err := c.Do("SELECT", cfg.Db); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		MaxIdle:         cfg.MaxIdle,
		MaxActive:       cfg.MaxActive,
		IdleTimeout:     time.Duration(cfg.IdleTimeout) * time.Second,
		Wait:            cfg.Wait,
		MaxConnLifetime: time.Duration(cfg.MaxConnLifetime) * time.Second,
	}

	c := R.Get()
	err := c.Err()
	if err != nil {
		return err
	}

	return c.Close()
}
