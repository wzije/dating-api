package src

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Json(ctx *fiber.Ctx, response Response) error {
	return ctx.Status(response.Code).JSON(response)
}

func JsonError(ctx *fiber.Ctx, code int, err error) error {
	return ctx.Status(code).JSON(Response{
		Code:    code,
		Message: err.Error(),
	})
}

func JsonErrorWithReason(ctx *fiber.Ctx, code int, message string, reason interface{}) error {
	return ctx.Status(code).JSON(Response{
		Code:    code,
		Message: message,
		Data:    reason,
	})
}
