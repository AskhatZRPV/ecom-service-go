package delete_deleteproduct

import (
	"ecomsvc/internal/application/product/deleteproduct"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase deleteproduct.UseCase
}

func New(usecase deleteproduct.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/product/:id"
}

func (h *handler) Method() string {
	return http.MethodDelete
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
