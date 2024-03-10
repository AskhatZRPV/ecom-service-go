package get_getallproducts

import (
	"ecomsvc/internal/application/auth/login"
	"ecomsvc/internal/interface/http_rest/common"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *handler) Handle(c *fiber.Ctx) error {
	var query queryParams
	if err := c.QueryParser(&query); err != nil {
		return errors.Wrap(err, "failed to parse login body")
	}

	r, err := h.usecase.Execute(c.Context(), query.toUsecasePayload())
	if err != nil {
		return h.resolveErr(c, err)
	}

	resp := responseFromResult(r)
	return c.JSON(resp)
}

func (h *handler) resolveErr(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, login.ErrAccountDoesNotExist):
		return common.ErrorBuilder(err).Detail("username", "Please sign up").Build()
	case errors.Is(err, login.ErrIncorrectPassword):
		return common.ErrorBuilder(err).Detail("password", "Incorrect password").Build()
	}

	return err
}
