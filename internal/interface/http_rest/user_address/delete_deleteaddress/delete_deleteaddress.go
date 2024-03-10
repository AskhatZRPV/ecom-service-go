package delete_deleteaddress

import (
	"ecomsvc/internal/application/address/deleteaddress"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase deleteaddress.UseCase
}

func New(usecase deleteaddress.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/address/:id"
}

func (h *handler) Method() string {
	return http.MethodDelete
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
