package usecase

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

// GetJournalByJournalID
func (u *usecase) GetJournalByJournalID(ctx context.Context, ID string) (*model.Journal, error) {
	// Retrieve data from DB
	journal, err := u.repo.GetJournalByJournalID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return journal, nil
}
