package quoteRepository

import (
	"context"
	"github.com/google/uuid"
)

func (r *QuoteRepository) DeleteQuote(id uuid.UUID) error {
	sqlQuery := `DELETE FROM "quotes" WHERE "ID" = $1`
	_, err := r.db.Exec(context.Background(), sqlQuery, id)
	if err != nil {
		return err
	}
	return nil
}
