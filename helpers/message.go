package helpers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type resMessage struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func MsgResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Code:    fiber.StatusOK,
		Message: msg,
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(resPonse)
}

func CrudResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Code:    fiber.StatusOK,
		Message: fmt.Sprintf(" %s data succesfully", msg),
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(resPonse)
}

func BadResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Code:    fiber.StatusBadRequest,
		Message: msg,
		Data:    data,
	}
	return c.Status(fiber.StatusBadRequest).JSON(resPonse)
}

func ServerResponse(c *fiber.Ctx, msg string, data interface{}) error {
	resPonse := &resMessage{
		Code:    fiber.StatusInternalServerError,
		Message: msg,
		Data:    data,
	}
	return c.Status(fiber.StatusInternalServerError).JSON(resPonse)
}

func NotFoundResponse(c *fiber.Ctx, data interface{}) error {
	resPonse := &resMessage{
		Code:    fiber.StatusNotFound,
		Message: "Not Found",
		Data:    data,
	}
	return c.Status(fiber.StatusNotFound).JSON(resPonse)
}