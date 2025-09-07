package otp

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"otp-auth/internal/db"
)

type Service struct {
	Redis *db.Redis
	TTL   time.Duration
}

func NewService(r *db.Redis, ttl time.Duration) *Service {
	return &Service{Redis: r, TTL: ttl}
}

func (s *Service) Generate(ctx context.Context, identifier string) (string, error) {
	code, err := randomDigits(6)
	if err != nil {
		return "", err
	}
	if err := s.Redis.SetOTP(ctx, identifier, code, s.TTL); err != nil {
		return "", err
	}
	return code, nil
}

func (s *Service) Verify(ctx context.Context, identifier, code string) (bool, error) {
	stored, err := s.Redis.GetOTP(ctx, identifier)
	if err != nil {
		return false, err
	}
	if stored == code {
		_ = s.Redis.DelOTP(ctx, identifier)
		return true, nil
	}
	return false, nil
}

func randomDigits(n int) (string, error) {
	max := big.NewInt(10)
	code := ""
	for i := 0; i < n; i++ {
		v, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		code += fmt.Sprint(v.Int64())
	}
	return code, nil
}
