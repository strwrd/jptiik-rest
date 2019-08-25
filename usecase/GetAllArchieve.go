package usecase

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

// GetAllArchieve
func (u *usecase) GetAllArchieve(ctx context.Context) ([]*model.Archieve, error) {
	// Retrieve data from DB
	archieves, err := u.repo.GetAllArchieve(ctx)
	if err != nil {
		return nil, err
	}

	return archieves, nil
}
