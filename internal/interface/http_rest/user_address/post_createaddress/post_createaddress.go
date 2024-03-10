package post_createaddress

import (
	"ecomsvc/internal/application/address/createaddress"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase createaddress.UseCase
}

func New(usecase createaddress.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/address"
}

func (h *handler) Method() string {
	return http.MethodPost
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
