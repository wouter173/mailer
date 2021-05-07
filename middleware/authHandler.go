package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wouter173/mailer/structs"
	"github.com/wouter173/mailer/utils"
)

func AuthHandler(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	if header == "" {
		c.Status(401)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Unauthorized"))
	}

	s := strings.Split(header, " ")

	if len(s) != 2 {
		c.Status(401)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Unauthorized"))
	}

	if s[0] != "Bearer" {
		c.Status(401)
		c.Set("Content-Type", "application/json")
		return c.Send(structs.NewResult("Unauthorized"))
	}

	for _, key := range utils.Readkeys().Keys {
		if key.Token == s[1] {
			c.Locals("Service", key.Service)
			return c.Next()
		}
	}

	c.Status(401)
	c.Set("Content-Type", "application/json")
	return c.Send(structs.NewResult("Unauthorized"))
}
