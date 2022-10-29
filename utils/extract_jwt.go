package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

type TokenMetadata struct {
	UserID    string
	SessionID string
	Expires   int64
}

func extractToken(ctx *fiber.Ctx) string {
	token := ctx.Get("Authorization")

	tokenSplit := strings.Split(token, " ")
	if len(tokenSplit) == 2 && tokenSplit[0] == "Bearer" {
		return tokenSplit[1]
	}

	return ""
}

func keyfunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
}

func verifyToken(ctx *fiber.Ctx) (*jwt.Token, error) {
	extractedToken := extractToken(ctx)

	token, err := jwt.Parse(extractedToken, keyfunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractJWT(ctx *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expires := int64(claims["expires"].(float64))
		userId := claims["userId"].(string)
		sessionId := claims["sessionId"].(string)

		return &TokenMetadata{
			UserID:    userId,
			SessionID: sessionId,
			Expires:   expires,
		}, nil
	}

	return nil, err
}
