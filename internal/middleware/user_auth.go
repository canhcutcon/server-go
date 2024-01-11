package middleware

import "github.com/golang-jwt/jwt"

type UserAuth struct {
	Username           string `json:"username"`
	jwt.StandardClaims        // This is the JWT standard claims object that already has the fields for expiration time, etc.
}
