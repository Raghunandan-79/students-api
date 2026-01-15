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

	"github.com/Raghunandan-79/students-api/internal/config"
	"github.com/Raghunandan-79/students-api/internal/http/handlers/student"
	"github.com/Raghunandan-79/students-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	router.HandleFunc("PUT /api/students/update/{id}", student.UpdateById(storage))
	router.HandleFunc("DELETE /api/students/deleteById/{id}", student.DeleteById(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started at", slog.String("address", cfg.Addr))

	// graceful shutdown
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-done

	// stopping server

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err1 := server.Shutdown(ctx)
	if err1 != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err1.Error()))
	}

	slog.Info("Server shutdown successfully")
}
