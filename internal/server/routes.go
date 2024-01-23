package server

import (

	"github.com/gofiber/fiber/v2"
)

func checkHealth(ctx *fiber.Ctx) error {
	status := `json:{"status" : "ok"}`
	return ctx.Status(fiber.StatusFound).JSON(status)
}
