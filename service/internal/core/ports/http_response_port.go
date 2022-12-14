package ports

import "github.com/gofiber/fiber/v2"

type IResponse interface {
	Json(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool) error
	JsonAuth(h *fiber.Ctx, httpCode int, message string, result interface{}, status bool, isValidAuthorize bool) error
	ErrorRequestBody(h *fiber.Ctx) error
}
