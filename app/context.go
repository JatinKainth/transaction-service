package app

import (
	"database/sql"
	"log/slog"

	"transaction_service/handler"
	"transaction_service/repository"
	"transaction_service/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AppContext struct {
	db     *sql.DB
	logger *slog.Logger
	echo   *echo.Echo

	transactionHandler *handler.TransactionHandler
}

func NewAppContext(db *sql.DB, logger *slog.Logger) *AppContext {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	transactionRepo := repository.NewTransactionRepository(db, logger)
	transactionService := service.NewTransactionService(transactionRepo, logger)
	transactionHandler := handler.NewTransactionHandler(transactionService, logger)

	return &AppContext{
		db:                 db,
		logger:             logger,
		echo:               e,
		transactionHandler: transactionHandler,
	}
}

func (app *AppContext) Run(addr string) error {
	transactionRouter := app.echo.Group("/transactionservice")

	transactionRouter.PUT("/transaction/:transaction_id", app.transactionHandler.PutTransaction)
	transactionRouter.GET("/transaction/:transaction_id", app.transactionHandler.GetTransaction)
	transactionRouter.GET("/types/:type", app.transactionHandler.GetTransactionsByType)
	transactionRouter.GET("/sum/:transaction_id", app.transactionHandler.GetTransactionSum)

	app.logger.Info("starting server", "addr", addr)
	return app.echo.Start(addr)
}
