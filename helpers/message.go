package helpers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type resMessage struct {
	Status    string         `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func MsgResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Status:    "Success",
		Message: msg,
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(resPonse)
}

func CrudResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Status:    "Success",
		Message: fmt.Sprintf(" %s data succesfully", msg),
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(resPonse)
}

func BadResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Status:    "Bad Request",
		Message: msg,
		Data:    data,
	}
	return c.Status(fiber.StatusBadRequest).JSON(resPonse)
}

func ServerResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Status:    "internal Server Error",
		Message: msg,
		Data:    data,
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resPonse)
}

func NotFoundResponse(c *fiber.Ctx, data interface{}) error {
	resPonse := &resMessage{
		Status:    "Not found",
		Message: "Document not found",
		Data:    data,
	}
	return c.Status(fiber.StatusNotFound).JSON(resPonse)
}