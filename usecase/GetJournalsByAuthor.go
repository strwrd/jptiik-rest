package usecase

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

func (u *usecase) GetJournalsByAuthor(ctx context.Context, author string) ([]*model.Journal, error) {
	return u.repo.GetJournalsByAuthor(ctx, author)
}
