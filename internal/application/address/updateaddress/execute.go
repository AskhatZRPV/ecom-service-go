package updateaddress

import (
	"context"
	"ecomsvc/internal/domain/useraddress"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	if _, err := i.uaRepo.UpdateById(ctx, useraddress.Update(p.FirstName, p.LastName, p.Address, p.City, p.PostalCode, p.Country, p.PhoneNumber)); err != nil {
		return errors.Wrap(err, "failed to persist refresh token")
	}
	return nil
}
