package get_getallproducts

import (
	"ecomsvc/internal/application/product/getallproducts"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase getallproducts.UseCase
}

func New(usecase getallproducts.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/product"
}

func (h *handler) Method() string {
	return http.MethodGet
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
