package quoteRepository

import "github.com/jackc/pgx/v5/pgxpool"

type QuoteRepository struct {
	db *pgxpool.Pool
}

func NewQuoteRepository(db *pgxpool.Pool) *QuoteRepository {
	return &QuoteRepository{
		db: db,
	}
}
