package usecase

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

// GetArchieveByCode
func (u *usecase) GetArchieveByCode(ctx context.Context, code string) (*model.Archieve, error) {
	// Retrieve data from DB
	archieve, err := u.repo.GetArchieveByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	// Retrieve data from DB
	journals, err := u.repo.GetJournalsByArchieveID(ctx, archieve.ID)
	if err != nil {
		return nil, err
	}

	// Append journals into archieve.journal object
	archieve.Journals = journals

	return archieve, nil
}
