package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nedokyrill/quotes-service/pkg/logger"
	"net/http"
	"os"
	"time"
)

type APIServer struct {
	httpServer *http.Server
}

func NewAPIServer(mux *mux.Router) *APIServer {
	return &APIServer{
		httpServer: &http.Server{
			Addr:         ":" + os.Getenv("API_PORT"),
			Handler:      mux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *APIServer) Start() {
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Logger.Fatal(err)
	}
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	switch ctx.Err() {
	case context.DeadlineExceeded:
		logger.Logger.Error("Timeout shutting down server")
	case nil:
		logger.Logger.Info("Shutdown completed before timeout.")
	default:
		logger.Logger.Error("Shutdown ended with:", ctx.Err())
	}

	return nil
}
