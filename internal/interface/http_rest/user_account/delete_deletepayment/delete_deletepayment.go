package delete_deletepayment

import (
	"ecomsvc/internal/application/account/deleteaccount"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase deleteaccount.UseCase
}

func New(usecase deleteaccount.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/payment/:id"
}

func (h *handler) Method() string {
	return http.MethodDelete
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
