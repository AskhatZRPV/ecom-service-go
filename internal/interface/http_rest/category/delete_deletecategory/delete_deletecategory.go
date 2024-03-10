package delete_deletecategory

import (
	"ecomsvc/internal/application/category/deletecategory"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase deletecategory.UseCase
}

func New(usecase deletecategory.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/category/:id"
}

func (h *handler) Method() string {
	return http.MethodDelete
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
