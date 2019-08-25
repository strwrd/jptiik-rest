package usecase

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
	"github.com/strwrd/jptiik-rest/repository/mysql"
)

// Usecase : usecase interface contract
type Usecase interface {
	GetAllArchieve(ctx context.Context) ([]*model.Archieve, error)
	GetAllJournal(ctx context.Context) ([]*model.Journal, error)
	GetArchieveByArchieveID(ctx context.Context, ID string) (*model.Archieve, error)
	GetArchieveByCode(ctx context.Context, code string) (*model.Archieve, error)
	GetJournalsByArchieveID(ctx context.Context, ID string) ([]*model.Journal, error)
	GetJournalByJournalID(ctx context.Context, ID string) (*model.Journal, error)
	GetJournalsByTitle(ctx context.Context, title string) ([]*model.Journal, error)
	GetJournalsByAuthor(ctx context.Context, title string) ([]*model.Journal, error)
}

type usecase struct {
	repo mysql.Repository
}

// NewUsecase : create usecase object
func NewUsecase(repo mysql.Repository) Usecase {
	return &usecase{
		repo,
	}
}
