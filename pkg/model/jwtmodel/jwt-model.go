package jwtmodel

import "github.com/synzofficial/nsd-synz-sharelib-api/pkg/enum"

type JwtContextKey struct{}
type JwtErrorKey struct{}

const (
	CUSTOM_CLAIMS_KEY = "custom_claims"
)

type CustomClaims struct {
	LogtoId  string        `json:"logto_id"`
	UserType enum.UserType `json:"user_type"`
	Email    string        `json:"email"`
	Phone    string        `json:"phone"`
}

func NewCustomClaims() CustomClaims {
	return CustomClaims{}
}
