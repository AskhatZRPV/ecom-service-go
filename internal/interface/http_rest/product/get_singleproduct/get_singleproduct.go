package get_singleproduct

import (
	"ecomsvc/internal/application/product/getproductbyid"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase getproductbyid.UseCase
}

func New(usecase getproductbyid.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/products/:id"
}

func (h *handler) Method() string {
	return http.MethodGet
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
