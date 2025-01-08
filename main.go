package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type app struct {
	logger *slog.Logger
}

func newApp(logger *slog.Logger) *app {
	return &app{
		logger: logger,
	}
}

func main() {
	// _, cancel := signal.NotifyContext(ctx, os.Interrupt)
	// defer cancel()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	addr := ":3000"

	app := newApp(logger)

	httpServer := &http.Server{
		Addr:    addr,
		Handler: app.routes(),
	}

	logger.Info(fmt.Sprintf("starting server. listening on %s", httpServer.Addr))
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}
