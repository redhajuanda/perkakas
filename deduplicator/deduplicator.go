package deduplicator

import (
	"time"

	"github.com/go-redis/redis"
)

const DefaultRedisExpireTime = 24

type Deduplicator struct {
	redis *redis.Client
}

func New(redis *redis.Client) *Deduplicator {
	return &Deduplicator{redis}
}

func (d *Deduplicator) Visit(key string) error {
	err := d.redis.Set(key, true, DefaultRedisExpireTime*time.Hour).Err()
	return err
}

func (d *Deduplicator) IsVisited(key string) bool {
	_, err := d.redis.Get(key).Result()
	if err != nil {
		return false
	}
	return true
}
