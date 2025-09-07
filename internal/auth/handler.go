package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"otp-auth/internal/db"
	"otp-auth/internal/otp"
)

type Handler struct {
	PG        *db.Postgres
	OTP       *otp.Service
	JWTSecret string
}

// RequestOTP godoc
// @Summary Request OTP
// @Description Send OTP code to email or phone number
// @Tags auth
// @Accept json
// @Produce json
// @Param data body auth.RequestOTPDTO true "Identifier"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /auth/request-otp [post]
func (h *Handler) RequestOTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Identifier string `json:"identifier"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad json", 400)
		return
	}

	ctx := context.Background()
	user, err := h.PG.FindOrCreateUser(ctx, strings.ToLower(body.Identifier))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	code, err := h.OTP.Generate(ctx, user.Identifier)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// TODO: replace with email/SMS sender
	// For now just log
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "code_dev": code})
}

// VerifyOTP godoc
// @Summary Verify OTP
// @Description Verify OTP code and issue JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param data body auth.VerifyOTPDTO true "Identifier + Code"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /auth/verify-otp [post]
func (h *Handler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Identifier string `json:"identifier"`
		Code       string `json:"code"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "bad json", 400)
		return
	}

	ctx := context.Background()
	ok, err := h.OTP.Verify(ctx, strings.ToLower(body.Identifier), body.Code)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if !ok {
		http.Error(w, "invalid code", 400)
		return
	}

	user, err := h.PG.FindOrCreateUser(ctx, strings.ToLower(body.Identifier))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_ = h.PG.UpdateVerified(ctx, user.ID)

	token, err := GenerateJWT(h.JWTSecret, user.ID, 7*24*time.Hour)
	if err != nil {
		http.Error(w, "token error", 500)
		return
	}

	_ = json.NewEncoder(w).Encode(map[string]any{
		"token": token,
		"user":  user,
	})
}
