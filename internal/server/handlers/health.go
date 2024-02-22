package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func CheckHealth(ctx *fiber.Ctx) error {
	status := `json:{"status" : "ok"}`
	return ctx.Status(fiber.StatusFound).JSON(status)
}
