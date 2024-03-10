package put_updatecart

import (
	"ecomsvc/internal/application/cart/updatecartitemquantity"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase updatecartitemquantity.UseCase
}

func New(usecase updatecartitemquantity.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/products"
}

func (h *handler) Method() string {
	return http.MethodPut
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
