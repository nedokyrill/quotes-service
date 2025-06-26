package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/nedokyrill/quotes-service/internal/server"
	"github.com/nedokyrill/quotes-service/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	// Init logger
	logger.InitLogger()
	defer logger.Logger.Sync()

	//// Load environment
	//err := godotenv.Load()
	//if err != nil {
	//	logger.Logger.Fatal("Error loading .env file")
	//}
	//
	//// Init postgres
	//conn, err := pgxpool.New(context.Background(), os.Getenv("POSTGRESQL_URL"))
	//if err != nil {
	//	logger.Logger.Fatal("Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//logger.Logger.Info("Connected to database successfully")
	//defer conn.Close()

	//// Init repositories
	//quotesRepo := quoteRepository.NewQuoteRepository(conn)
	//
	//// Init services
	//quotesServ := quotesService.NewQuotesService(quotesRepo)
	//
	//// Init handlers
	//quotesHand := quoteHandlers.NewQuoteHandler(quotesServ)

	// Init default route
	router := mux.NewRouter()
	router.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods("GET")

	// Register routes
	//quotesHand.RegisterRoutes(router)

	// Init n conf API server
	srv := server.NewAPIServer(router)

	// Start server
	go srv.Start()
	logger.Logger.Info("API server started successfully on port " + os.Getenv("API_PORT"))

	// Shutdown server (with graceful shutdown)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Logger.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatalw("Shutdown error",
			"error", err)
	}
}
