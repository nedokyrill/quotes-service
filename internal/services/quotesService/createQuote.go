package quotesService

import (
	"encoding/json"
	"github.com/nedokyrill/quotes-service/internal/models/quoteModel"
	"github.com/nedokyrill/quotes-service/pkg/logger"
	"net/http"
)

func (s *QuotesService) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var newQuote struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}

	err := json.NewDecoder(r.Body).Decode(&newQuote)
	if err != nil {
		logger.Logger.Error("Error with parsing JSON: ", err)
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.quoteRepo.CreateQuote(quoteModel.Quote{Author: newQuote.Author, Text: newQuote.Quote})
	if err != nil {
		logger.Logger.Error("Failed to create quote:", err)
		http.Error(w, "Failed to create quote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	logger.Logger.Info("Quote created successfully")
}
