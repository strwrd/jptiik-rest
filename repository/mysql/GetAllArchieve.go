package mysql

import (
	"context"

	"github.com/strwrd/jptiik-rest/model"
)

// GetAllArchieve : return all archieve in DB
func (r *repository) GetAllArchieve(ctx context.Context) ([]*model.Archieve, error) {
	// Mysql Query
	query := "SELECT a.archieve_id, a.code, a.link, a.published, a.created_at, a.updated_at FROM archieve AS a ORDER BY a.published DESC"

	// Creating archieve array(slice)
	archieves := make([]*model.Archieve, 0)

	// Execute query
	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterating result from execution
	for rows.Next() {
		// Create archieve Object
		archieve := new(model.Archieve)

		// Copy data from DB into archieve object
		err := rows.Scan(
			&archieve.ID,
			&archieve.Code,
			&archieve.Link,
			&archieve.Published,
			&archieve.CreatedAt,
			&archieve.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Insert archieve into archieve array(slice)
		archieves = append(archieves, archieve)
	}

	// Check if any error during execution
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return archieves, nil
}
