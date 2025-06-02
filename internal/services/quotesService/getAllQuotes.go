package quotesService

import (
	"encoding/json"
	"github.com/nedokyrill/quotes-service/pkg/logger"
	"net/http"
	"strings"
)

func (s *QuotesService) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	author := strings.TrimSpace(r.URL.Query().Get("author"))

	quotes, err := s.quoteRepo.GetAllQuotes(author)
	if err != nil {
		logger.Logger.Error("Error while getting all quotes: ", err)
		http.Error(w, "Failed to fetch quotes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(quotes); err != nil {
		logger.Logger.Error("Error while encoding quotes: ", err)
		http.Error(w, "Failed to encode quotes", http.StatusInternalServerError)
	}

	logger.Logger.Info("Successfully fetched all quotes")
}
