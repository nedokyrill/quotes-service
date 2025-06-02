package quotesService

import "github.com/nedokyrill/quotes-service/internal/repository"

type QuotesService struct {
	quoteRepo repository.QuoteRepositoryInterface
}

func NewQuotesService(QuoteRepository repository.QuoteRepositoryInterface) *QuotesService {
	return &QuotesService{
		quoteRepo: QuoteRepository,
	}
}
