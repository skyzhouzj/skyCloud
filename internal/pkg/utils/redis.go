package utils

import (
	"github.com/go-redis/redis/v7"
	"github.com/mojocn/base64Captcha"
	"time"
)

const CAPTCHA = "captcha:"

type redisStore struct {
	base64Captcha.Store
	redis *redis.Client
}

var RedisStore = new(redisStore)

func (r redisStore) GetStore(redis *redis.Client) redisStore {
	r.redis = redis
	return r
}

//set a capt
func (r redisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := r.redis.Set(key, value, time.Minute*2).Err()
	if err != nil {
		return err
	} else {
		return nil
	}
}

//get a capt
func (r redisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := r.redis.Get(key).Result()
	if err != nil {
		return ""
	}
	if clear {
		err := r.redis.Del(key).Err()
		if err != nil {
			return ""
		}
	}
	return val
}

//verify a capt
func (r redisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
