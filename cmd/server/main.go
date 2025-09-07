package main

import (
	"log"
	"net/http"
	"time"

	"otp-auth/internal/auth"
	"otp-auth/internal/config"
	"otp-auth/internal/db"
	"otp-auth/internal/otp"
	"otp-auth/internal/server"

	_ "otp-auth/internal/docs"
)

func main() {
	cfg := config.Load()

	// Postgres
	pg, err := db.NewPostgres(cfg.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}

	// Redis
	rd, err := db.NewRedis(cfg.RedisURL)
	if err != nil {
		log.Fatal(err)
	}

	otpService := otp.NewService(rd, 2*time.Minute)

	h := &auth.Handler{
		PG:        pg,
		OTP:       otpService,
		JWTSecret: cfg.JWTSecret,
	}

	srv := server.New(h)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", srv))
}
