package quoteHandlers

import (
	"github.com/gorilla/mux"
	"github.com/nedokyrill/quotes-service/internal/services"
)

type QuoteHandler struct {
	QuoteService services.QuoteServiceInterface
}

func NewQuoteHandler(quoteService services.QuoteServiceInterface) *QuoteHandler {
	return &QuoteHandler{
		QuoteService: quoteService,
	}
}

func (h *QuoteHandler) RegisterRoutes(router *mux.Router) {
	quoteRouter := router.PathPrefix("/quotes").Subrouter()

	quoteRouter.HandleFunc("", h.QuoteService.CreateQuote).Methods("POST")
	quoteRouter.HandleFunc("/random", h.QuoteService.GetRandomQuote).Methods("GET")
	quoteRouter.HandleFunc("", h.QuoteService.GetAllQuotes).Methods("GET")
	quoteRouter.HandleFunc("/{id}", h.QuoteService.DeleteQuote).Methods("DELETE")

}
