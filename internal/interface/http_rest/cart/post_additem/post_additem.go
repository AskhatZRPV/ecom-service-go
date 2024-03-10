package post_additem

import (
	addproducttocart "ecomsvc/internal/application/cart/addproductocart"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase addproducttocart.UseCase
}

func New(usecase addproducttocart.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/cart"
}

func (h *handler) Method() string {
	return http.MethodPost
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
