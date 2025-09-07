package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(url string) (*Redis, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return &Redis{Client: redis.NewClient(opt)}, nil
}

func (r *Redis) SetOTP(ctx context.Context, key, code string, ttl time.Duration) error {
	return r.Client.Set(ctx, "otp:"+key, code, ttl).Err()
}

func (r *Redis) GetOTP(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, "otp:"+key).Result()
}

func (r *Redis) DelOTP(ctx context.Context, key string) error {
	return r.Client.Del(ctx, "otp:"+key).Err()
}

func (r *Redis) Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	cnt, err := r.Client.Incr(ctx, "rate:"+key).Result()
	if err != nil {
		return false, err
	}
	if cnt == 1 {
		_ = r.Client.Expire(ctx, "rate:"+key, window).Err()
	}
	return int(cnt) <= limit, nil
}
