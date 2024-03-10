package put_updatecategory

import (
	"ecomsvc/internal/application/category/updatecategory"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase updatecategory.UseCase
}

func New(usecase updatecategory.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/category"
}

func (h *handler) Method() string {
	return http.MethodPut
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
