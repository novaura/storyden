package account

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Southclaws/storyden/internal/errtag"
	"github.com/Southclaws/storyden/pkg/resources/account"
	"github.com/Southclaws/storyden/pkg/services/authentication"
)

func (s *service) Get(ctx context.Context, id account.AccountID) (*account.Account, error) {
	subject, err := authentication.GetAccountID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account")
	}

	if subject != id {
		return nil, errtag.Wrap(errors.New("not owner"), errtag.PermissionDenied{})
	}

	acc, err := s.account_repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account by ID")
	}

	return acc, nil
}
