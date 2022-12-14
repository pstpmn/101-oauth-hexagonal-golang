package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	ports2 "learn-oauth2/internal/core/ports"
)

type HTTPMiddleware struct {
	membersUseCase   ports2.MembersUseCase
	response         ports2.IResponse
	authorizationKey string
}

func (H HTTPMiddleware) Authorize(h *fiber.Ctx) error {
	clientToken := h.Get("Authorize-token")
	e, err := H.membersUseCase.Authorization(clientToken, H.authorizationKey)
	if err != nil {
		return H.response.JsonAuth(h, fiber.StatusBadRequest, fmt.Sprintf("%s", err), nil, false, false)
	}

	h.Locals("memberId", fmt.Sprintf("%s", e["memberId"]))
	h.Locals("username", fmt.Sprintf("%s", e["username"]))
	h.Locals("firstName", fmt.Sprintf("%s", e["firstName"]))
	h.Locals("lastName", fmt.Sprintf("%s", e["lastName"]))
	h.Locals("dob", fmt.Sprintf("%s", e["dob"]))
	h.Locals("createdAt", fmt.Sprintf("%s", e["createdAt"]))

	return h.Next()
}

func NewHTTPMiddleware(membersUseCase ports2.MembersUseCase, response ports2.IResponse, key string) ports2.IMiddlewares {
	return &HTTPMiddleware{membersUseCase: membersUseCase, response: response, authorizationKey: key}
}
