package repository

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/quotes-service/internal/models/quoteModel"
)

type QuoteRepositoryInterface interface {
	GetRandomQuote() (*quoteModel.Quote, error)
	GetAllQuotes(author string) (*[]quoteModel.Quote, error)
	CreateQuote(quoteModel.Quote) error
	DeleteQuote(id uuid.UUID) error
}
