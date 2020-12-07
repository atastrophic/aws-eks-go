package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// User represents a user
type User struct {
	ID       uuid.UUID `json:"ID"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"-" validate:"required"`
}

// Session ...
type Session struct {
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token"`
}

// Claims ...
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
