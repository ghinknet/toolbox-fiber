package ip

import (
	"github.com/gofiber/fiber/v3"
	"go.gh.ink/toolbox/expr"
)

// GetIP returns accurate IP address of visitor
func GetIP(c fiber.Ctx) string {
	return expr.Ternary(len(c.IPs()) > 0, c.IPs(), []string{c.IP()})[0]
}
