package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		log.Info("request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("duration", duration),
			zap.Any("request_id", c.Locals("request_id")),
		)

		return err
	}
}
