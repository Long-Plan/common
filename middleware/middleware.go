package middleware

import (
	"github.com/Long-Plan/common/logger"
	"github.com/gofiber/fiber/v2"
)

type middleware struct {
	jwtPublicKey []byte
	issuer       string
	audience     string
	logger       logger.ILogger
}

func NewMiddleware(jwtPublicKey []byte, issuer string, audience string, logger logger.ILogger) IMiddleware {
	return &middleware{
		logger: logger,
	}
}

type IMiddleware interface {
	AuthMiddleware(c *fiber.Ctx) error
}
