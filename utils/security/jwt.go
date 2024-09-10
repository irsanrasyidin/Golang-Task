package security

import (
	"fmt"
	"time"
	"usecase-1/config"
	"usecase-1/model"
	"usecase-1/utils/exceptions"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user model.Usecase2LoginModel) (string, error) {
	cfg, err := config.NewConfig()
	exceptions.CheckError(err)

	now := time.Now().UTC()
	end := now.Add(cfg.AccessTokenLifeTime)

	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.ApplicationName,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(cfg.JwtSignatureKey)
	fmt.Printf("%v %v", ss, err)
	if err != nil {
		return "", fmt.Errorf("failed to create access token: %v", err)
	}
	return ss, nil
}

func VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	cfg, _ := config.NewConfig()

	// Parsing JWT Token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok || method != (*jwt.SigningMethodHMAC)(cfg.JwtSigningMethod) {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Memastikan klaim token adalah tipe MapClaims dan token valid
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Periksa apakah 'iss' dalam klaim sesuai dengan ApplicationName yang diharapkan
	if claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid issuer in token")
	}

	return claims, nil
}
