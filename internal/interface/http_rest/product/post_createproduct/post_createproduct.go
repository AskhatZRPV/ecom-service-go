package post_createproduct

import (
	"ecomsvc/internal/application/product/createproduct"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase createproduct.UseCase
}

func New(usecase createproduct.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/product"
}

func (h *handler) Method() string {
	return http.MethodPost
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
