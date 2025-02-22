package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Long-Plan/common/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func (m middleware) AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		m.logger.Error("missing token", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("missing token")))
	}

	jwtToken, err := extractToken(tokenString)
	if err != nil {
		m.logger.Error("missing token header", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("missing token header")))
	}

	decodedToken, err := decodeJWTToken(jwtToken)
	if err != nil {
		m.logger.Error("missing token", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.BadRequestCode, err))
	}

	userUUID, ok := decodedToken["user_uuid"]
	if !ok {
		m.logger.Error("missing user uuid", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("missing user uuid")))
	}

	userRole, ok := decodedToken["user_role"]
	if !ok {
		m.logger.Error("missing user role", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("missing user role")))
	}

	organizationUUID, ok := decodedToken["organization_uuid"]
	if !ok {
		m.logger.Error("missing organization uuid", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("missing organization uuid")))
	}

	claims := jwt.MapClaims{}
	jwt.WithIssuer(m.issuer)
	jwt.WithAudience(m.audience)

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.jwtPublicKey), nil
	})

	if err != nil {
		m.logger.Error(fmt.Sprintf("failed to parse token: %v", err), zap.String("tag", "require token middleware"))
		m.logger.Error("failed to parse token", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("failed to parse token")))
	}

	if token.Method.Alg() != "HS256" {
		m.logger.Error("invalid token algorithm", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("invalid token algorithm")))
	}

	if !token.Valid {
		m.logger.Error("invalid token", zap.String("tag", "require token middleware"))
		return response.WriteError(c, response.NewErrorData(response.UnauthorizedCode, errors.New("invalid token")))
	}

	c.Locals("user_uuid", userUUID)
	c.Locals("user_role", userRole)
	c.Locals("organization_uuid", organizationUUID)

	return c.Next()
}

func extractToken(rawToken string) (string, error) {
	if len(rawToken) > 7 && strings.HasPrefix(rawToken, "Bearer ") {
		return rawToken[7:], nil
	}
	return "", errors.New("token is required")
}

func decodeJWTToken(tokenString string) (map[string]interface{}, error) {

	// Split part
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	// Decode
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, errors.New("failed to decode token payload")
	}

	// Unmarshal
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, errors.New("failed to unmarshal token claims")
	}

	return claims, nil
}
