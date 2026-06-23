package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

type JWTClaim struct {
	UserID   uint   `json:"user_id"`
	RoleName string `json:"role"`
	AgencyID *uint  `json:"agency_id"`
	jwt.RegisteredClaims
}

func InitJWT() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "supersecretkey" // Default fallback
	}
	jwtKey = []byte(secret)
}

func GenerateToken(userID uint, roleName string, agencyID *uint) (string, error) {
	InitJWT() // ensure key is loaded
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaim{
		UserID:   userID,
		RoleName: roleName,
		AgencyID: agencyID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	InitJWT()
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
