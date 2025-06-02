package quoteRepository

import (
	"context"
	"github.com/google/uuid"
	"github.com/nedokyrill/quotes-service/internal/models/quoteModel"
)

func (r *QuoteRepository) GetAllQuotes(authorName string) (*[]quoteModel.Quote, error) {
	var quotes []quoteModel.Quote

	if authorName == "" {
		sqlQuery := `SELECT "ID", "author", "text" FROM quotes`

		rows, err := r.db.Query(context.Background(), sqlQuery)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var id uuid.UUID
			var author string
			var text string

			if err = rows.Scan(&id, &author, &text); err != nil {
				return nil, err
			}

			quotes = append(quotes, quoteModel.Quote{Id: id, Author: author, Text: text})
		}
	} else {
		sqlQuery := `SELECT "ID", "author", "text" FROM quotes WHERE "author" LIKE $1`

		rows, err := r.db.Query(context.Background(), sqlQuery, authorName)
		if err != nil {
			return nil, err
		}

		defer rows.Close()
		for rows.Next() {
			var id uuid.UUID
			var author string
			var text string

			if err = rows.Scan(&id, &author, &text); err != nil {
				return nil, err
			}

			quotes = append(quotes, quoteModel.Quote{Id: id, Author: author, Text: text})
		}
	}

	return &quotes, nil
}
