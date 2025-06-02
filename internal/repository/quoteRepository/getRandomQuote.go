package quoteRepository

import (
	"context"
	"github.com/nedokyrill/quotes-service/internal/models/quoteModel"
)

func (r *QuoteRepository) GetRandomQuote() (*quoteModel.Quote, error) {
	var quote quoteModel.Quote
	sqlQuery := `SELECT "ID", "author", "text" FROM "quotes" ORDER BY RANDOM() LIMIT 1;`
	err := r.db.QueryRow(context.Background(), sqlQuery).Scan(&quote.Id, &quote.Author, &quote.Text)
	if err != nil {
		return nil, err
	}
	return &quote, nil
}
