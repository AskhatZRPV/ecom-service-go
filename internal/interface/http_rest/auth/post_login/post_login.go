package post_login

import (
	"ecomsvc/internal/application/auth/login"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase login.UseCase
}

func New(usecase login.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/auth/login"
}

func (h *handler) Method() string {
	return http.MethodPost
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
