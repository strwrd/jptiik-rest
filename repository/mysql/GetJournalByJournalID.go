package mysql

import (
	"context"
	"database/sql"

	"github.com/strwrd/jptiik-rest/model"
)

// GetJournalByJournalID: return journal based on journalID
func (r *repository) GetJournalByJournalID(ctx context.Context, ID string) (*model.Journal, error) {
	// Mysql query
	query := "SELECT j.journal_id, j.archieve_id, j.title, j.authors, j.abstract, j.link, j.pdf_link, j.published, j.created_at, j.updated_at FROM journal AS j WHERE j.journal_id = ?"

	// Create journal object
	journal := new(model.Journal)

	// Query execution and copy result on journal object
	err := r.conn.QueryRowContext(ctx, query, ID).Scan(
		&journal.ID,
		&journal.ArchieveID,
		&journal.Title,
		&journal.Authors,
		&journal.Abstract,
		&journal.Link,
		&journal.PDFLink,
		&journal.Published,
		&journal.CreatedAt,
		&journal.UpdatedAt,
	)

	// Check if any error during execution
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrDataNotFound
		}
		return nil, err
	}

	return journal, nil
}
