package deletecart

import (
	"context"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	_, err := i.ssRepo.Delete(ctx, p.ID)

	// FIXME:
	if err != nil {
		return errors.Wrap(err, "cant delete shopping session with specified id")
	}
	return nil
}
