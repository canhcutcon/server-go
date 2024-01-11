package services

import (
	"server-go/internal/configs"
	"server-go/internal/pkg/log"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger log.Logger
	cfg    *configs.Config
}

type Token struct {
	UserID   uint
	Username string
	Phone    string
	Email    string
	Role     string
}

type TokenDetails struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	AccessTokenExp  string `json:"access_token_exp"`
	RefreshTokenExp string `json:"refresh_token_exp"`
}

func NewTokenService(cfg *configs.Config) *TokenService {
	logger := log.NewLogger(cfg)
	return &TokenService{
		logger: logger,
		cfg:    cfg,
	}
}

func (s *TokenService) GenerateToken(user *Token) (*TokenDetails, error) {
	td := TokenDetails{}
	// set the expiration time for the access token
	td.AccessTokenExp = time.Now().Add(time.Minute * time.Duration(s.cfg.Jwt.Expires)).Format(time.RFC3339) // 1 hour
	td.RefreshTokenExp = time.Now().Add(time.Minute * time.Duration(s.cfg.Jwt.Expires)).Format(time.RFC3339)

	// create the access token claims
	atc := jwt.MapClaims{} // access token claims
	atc["user_id"] = user.UserID
	atc["username"] = user.Username
	atc["phone"] = user.Phone
	atc["email"] = user.Email
	atc["role"] = user.Role
	atc["exp"] = td.AccessTokenExp

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc) // create access token with claims above

	var err error
	td.AccessToken, err = at.SignedString([]byte(s.cfg.Jwt.Secret)) // sign the access token with the secret key
	if err != nil {
		return nil, err
	}

	// create the refresh token claims
	rtc := jwt.MapClaims{} // refresh token claims
	rtc["user_id"] = user.UserID
	rtc["exp"] = td.RefreshTokenExp

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)             // create refresh token with claims above
	td.RefreshToken, err = rt.SignedString([]byte(s.cfg.Jwt.Secret)) // sign the refresh token with the secret key
	if err != nil {
		return nil, err
	}

	return &td, nil
}

func (s *TokenService) VerifyToken(tokenString string) (*jwt.Token, error) {
	// From token using the secret key to verify
	at, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // check if the signing method is HMAC
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.cfg.Jwt.Secret), nil // return the secret key
	})

	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) ExtractTokenMetadata(tokenString string) (*Token, error) {
	token, err := s.VerifyToken(tokenString) // verify the token
	if err != nil {
		return nil, err
	}

	claims, err := s.extractClaims(token) // extract the claims
	if err != nil {
		return nil, err
	}

	return &Token{
		UserID:   uint(claims["user_id"].(float64)), // extract the user_id
		Username: claims["username"].(string),       // extract the username
		Phone:    claims["phone"].(string),          // extract the phone
		Email:    claims["email"].(string),          // extract the email
		Role:     claims["role"].(string),           // extract the role
	}, nil
}

func (s *TokenService) extractClaims(token *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := token.Claims.(jwt.MapClaims) // the token claims should conform to MapClaims
	if !ok || !token.Valid {
		return nil, jwt.ValidationError{}
	}

	return claims, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = make(map[string]interface{})
	t, err := s.VerifyToken(token) // verify the token
	if err != nil {
		return
	}

	claims, ok := t.Claims.(jwt.MapClaims) // the token claims should conform to MapClaims
	if ok && t.Valid {
		for k, v := range claims { // extract the claims
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, err
}
