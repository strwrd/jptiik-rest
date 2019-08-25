package mysql

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

// GetAllJournal : return all journals in DB
func (r *repository) GetAllJournal(ctx context.Context) ([]*model.Journal, error) {
	// Mysql query
	query := "SELECT j.journal_id, j.archieve_id, j.title, j.authors, j.abstract, j.link, j.pdf_link, j.published, j.created_at, j.updated_at FROM journal AS j ORDER BY j.published DESC"

	// Create journal array(slice)
	journals := make([]*model.Journal, 0)

	// Execute query
	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterating result from execution
	for rows.Next() {
		// create journal object
		journal := new(model.Journal)

		// Copy data from DB into journal object
		err := rows.Scan(
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
		if err != nil {
			return nil, err
		}

		// Insert journal into journal array(slice)
		journals = append(journals, journal)
	}

	// Check if any error during execution
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return journals, nil
}
