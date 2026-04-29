package auth

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

// GetAuthorization returns authorisation from authorisation header
func GetAuthorization(c fiber.Ctx) (string, AuthorizationType) {
	authorization := c.Get("Authorization")

	if authorization == "" {
		return "", AuthorizationTypeNone
	}

	switch {
	case strings.HasPrefix(authorization, "Bearer "):
		return strings.TrimPrefix(authorization, "Bearer "), AuthorizationTypeBearer
	case strings.HasPrefix(authorization, "Basic "):
		return strings.TrimPrefix(authorization, "Basic "), AuthorizationTypeBasic
	}
	return authorization, AuthorizationTypeOther
}
