package mysql

import (
	"context"
	"database/sql"

	"github.com/strwrd/jptiik-rest/model"
)

// GetArchieveByArchieveID : return archieve based on archieveID
func (r *repository) GetArchieveByArchieveID(ctx context.Context, ID string) (*model.Archieve, error) {
	// Mysql query
	query := "SELECT a.archieve_id, a.code, a.link, a.published, a.created_at, a.updated_at FROM archieve AS a WHERE a.archieve_id = ?"

	// Creating archieve object
	archieve := new(model.Archieve)

	// Query execution then copy result into archieve object
	err := r.conn.QueryRowContext(ctx, query, ID).Scan(
		&archieve.ID,
		&archieve.Code,
		&archieve.Link,
		&archieve.Published,
		&archieve.CreatedAt,
		&archieve.UpdatedAt,
	)

	// Check if any error during execution
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrDataNotFound
		}
		return nil, err
	}

	return archieve, nil
}
