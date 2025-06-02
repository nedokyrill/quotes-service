package quotesService

import (
	"encoding/json"
	"github.com/nedokyrill/quotes-service/pkg/logger"
	"net/http"
)

func (s *QuotesService) GetRandomQuote(w http.ResponseWriter, _ *http.Request) {
	quote, err := s.quoteRepo.GetRandomQuote()
	if err != nil {
		logger.Logger.Error("Error while getting random quote: ", err)
		http.Error(w, "Failed to fetch random quote", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(quote); err != nil {
		logger.Logger.Error("Error while encoding random quote: ", err)
		http.Error(w, "Failed to encode quote", http.StatusInternalServerError)
	}

	logger.Logger.Info("Successfully fetched random quote")
}
