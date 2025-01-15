package main

import (
	"log/slog"
	"os"

	"transaction_service/app"
	"transaction_service/config"
	"transaction_service/pkg/db"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	db, err := db.Initialize(&cfg.Database)
	if err != nil {
		logger.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	appContext := app.NewAppContext(db, logger)
	logger.Info("application initialized, starting server")

	if err := appContext.Run(":8080"); err != nil {
		logger.Error("server error", "error", err)
		os.Exit(1)
	}
}
