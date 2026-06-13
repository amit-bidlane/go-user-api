package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go-user-api/internal/logger"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)

		logger.Log.Info("request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("requestID", c.Locals("requestID").(string)),
		)

		return err
	}
}