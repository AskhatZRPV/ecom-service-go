package createaddress

import (
	"context"
	"ecomsvc/internal/domain/useraddress"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (Result, error) {
	res, err := i.uaRepo.Save(ctx, useraddress.New(p.UserId, p.FirstName, p.LastName, p.Address, p.City, p.PostalCode, p.Country, p.PhoneNumber))
	if err != nil {
		return 0, errors.Wrap(err, "failed to persist refresh token")
	}
	return res, nil
}
