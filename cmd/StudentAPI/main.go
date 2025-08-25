package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/Sachiinkk/student-api/internal"
	"github.com/Sachiinkk/student-api/internal/http/handlers/student"
)



func main() {

	//config set up

	cfg := config.MustLoad()

	//Router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/student", student.New())

	// server setup

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}


	slog.Info("server started", slog.String("address",cfg.Addr))
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done

	slog.Info("Shutting Down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown  server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown succesfully")
}
