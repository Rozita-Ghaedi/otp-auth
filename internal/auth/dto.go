package auth

type RequestOTPDTO struct {
	Identifier string `json:"identifier"`
}

type VerifyOTPDTO struct {
	Identifier string `json:"identifier"`
	Code       string `json:"code"`
}
