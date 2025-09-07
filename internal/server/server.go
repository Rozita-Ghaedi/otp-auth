package server

import (
	"net/http"

	"otp-auth/internal/auth"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(handler *auth.Handler) http.Handler {
	r := chi.NewRouter()

	// Auth routes
	r.Post("/auth/request-otp", handler.RequestOTP)
	r.Post("/auth/verify-otp", handler.VerifyOTP)

	// Swagger docs
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
