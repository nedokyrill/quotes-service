package services

import "net/http"

type QuoteServiceInterface interface {
	CreateQuote(w http.ResponseWriter, r *http.Request)
	GetAllQuotes(w http.ResponseWriter, r *http.Request)
	GetRandomQuote(w http.ResponseWriter, r *http.Request)
	DeleteQuote(w http.ResponseWriter, r *http.Request)
}
