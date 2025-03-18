package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Connection struct {
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	Password string
	DB       int
}

func (c *Connection) String() *string {
	conn := fmt.Sprintf("%s:%s", c.Host, c.Port)
	return &conn
}

func (c *Connection) getRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     *c.String(),
		Password: c.Password,
		DB:       c.DB,
	}
}

func (c *Connection) Connect() *redis.Client {
	return redis.NewClient(c.getRedisOptions())
}
