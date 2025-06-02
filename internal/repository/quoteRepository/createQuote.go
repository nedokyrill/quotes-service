package quoteRepository

import (
	"context"
	"github.com/nedokyrill/quotes-service/internal/models/quoteModel"
)

func (r *QuoteRepository) CreateQuote(quote quoteModel.Quote) error {
	sqlQuery := `INSERT INTO quotes ("author", "text") VALUES ($1, $2)`
	_, err := r.db.Exec(context.Background(), sqlQuery, quote.Author, quote.Text)
	if err != nil {
		return err
	}
	return nil
}
