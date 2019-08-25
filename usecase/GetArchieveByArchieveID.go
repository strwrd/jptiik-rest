package usecase

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

// GetArchieveByArchieveID
func (u *usecase) GetArchieveByArchieveID(ctx context.Context, ID string) (*model.Archieve, error) {
	// Retrieve data from DB
	archieve, err := u.repo.GetArchieveByArchieveID(ctx, ID)
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
