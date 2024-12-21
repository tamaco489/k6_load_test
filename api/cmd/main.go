package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tamaco489/k6_load_test/api/jester/internal/controller"
)

func main() {
	ctx := context.Background()

	server, err := controller.NewHJesterAPIServer()
	if err != nil {
		panic(err)
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.ErrorContext(ctx, "Failed to listen and serve", "error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.WarnContext(ctx, "shutdown server...", "error", err)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Microsecond)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "shutdown server...", "error", err)
	}
	<-ctx.Done()
}
