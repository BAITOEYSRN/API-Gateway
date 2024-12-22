package models

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AppSuccessResponse struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

func ResponseSuccess(c *fiber.Ctx, data any) error {
	return c.Status(200).JSON(AppSuccessResponse{
		Status: true,
		Data:   data,
	})
}

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

type AppErrorResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error"`
}

func ResponseError(c *fiber.Ctx, code int, err error) error {
	return c.Status(code).JSON(AppErrorResponse{
		Status: false,
		Error:  err.Error(),
	})
}

// func ResponseError(c *fiber.Ctx, code int, err error) error {
// 	return c.Status(code).JSON(AppErrorResponse{
// 		Status: false,
// 		Error:  err.Error(),
// 	})
// }
