package quotesService

import (
	"github.com/google/uuid"
	"github.com/nedokyrill/quotes-service/pkg/logger"
	"net/http"
	"strings"
)

func (s *QuotesService) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	idStr := url[len(url)-1]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	err = s.quoteRepo.DeleteQuote(id)
	if err != nil {
		logger.Logger.Error("Error deleting quote:", err)
		http.Error(w, "Failed to delete quote", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	logger.Logger.Info("Deleted quote successfully")
}
