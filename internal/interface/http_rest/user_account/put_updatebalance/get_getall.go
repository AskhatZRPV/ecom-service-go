package put_updatebalance

import (
	"ecomsvc/internal/application/account/updateaccount"
	"ecomsvc/internal/interface/http_rest/common"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase updateaccount.UseCase
}

func New(usecase updateaccount.UseCase) common.Handler {
	return &handler{usecase}
}

func (h *handler) Pattern() string {
	return "/account"
}

func (h *handler) Method() string {
	return http.MethodPut
}

func (h *handler) Middleware() []fiber.Handler {
	return []fiber.Handler{}
}
