package utils

import (
	"beep-work-backend/config"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

type Tokens struct {
	Access  string
	Refresh string
}

func GenerateTokens(userId string, sessionId string) (*Tokens, error) {
	access, err := generateAccessToken(userId, sessionId)
	if err != nil {
		return nil, err
	}

	refresh, err := generateNewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func generateAccessToken(userId string, sessionId string) (string, error) {
	secret := config.GetEnv("JWT_ACCESS_SECRET")
	lifetime, _ := strconv.Atoi(config.GetEnv("JWT_ACCESS_LIFETIME"))

	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["sessionId"] = sessionId
	claims["expires"] = time.Now().Add(time.Minute * time.Duration(lifetime)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateNewRefreshToken() (string, error) {
	hash := sha256.New()
	refresh := config.GetEnv("JWT_REFRESH_SECRET") + time.Now().String()

	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}

	lifetime, _ := strconv.Atoi(config.GetEnv("JWT_REFRESH_LIFETIME"))
	expireDate := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(lifetime)).Unix())

	tokenString := hex.EncodeToString(hash.Sum(nil)) + "." + expireDate

	return tokenString, nil
}
