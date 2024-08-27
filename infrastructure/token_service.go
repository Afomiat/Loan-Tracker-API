package infrastructure

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

type TokenService struct {
    AccessSecret  string
    RefreshSecret string
    AccessExpiry  time.Duration
    RefreshExpiry time.Duration
}

func NewTokenService(accessSecret, refreshSecret string, accessExpiry, refreshExpiry time.Duration) *TokenService {
    return &TokenService{
        AccessSecret:  accessSecret,
        RefreshSecret: refreshSecret,
        AccessExpiry:  accessExpiry,
        RefreshExpiry: refreshExpiry,
    }
}

func (ts *TokenService) GenerateAccessToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(ts.AccessExpiry).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(ts.AccessSecret))
}

func (ts *TokenService) GenerateRefreshToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(ts.RefreshExpiry).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(ts.RefreshSecret))
}

func (ts *TokenService) ParseToken(tokenString string, secret string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
}
