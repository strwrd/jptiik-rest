package usecase

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

func (u *usecase) GetJournalsByTitle(ctx context.Context, title string) ([]*model.Journal, error) {
	return u.repo.GetJournalsByTitle(ctx, title)
}
