package handler

import (
	"log/slog"
	"net/http"

	"transaction_service/api"
	"transaction_service/pkg/util"
	"transaction_service/service"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	service *service.TransactionService
	logger  *slog.Logger
}

func NewTransactionHandler(service *service.TransactionService, logger *slog.Logger) *TransactionHandler {
	return &TransactionHandler{
		service: service,
		logger:  logger,
	}
}

func (h *TransactionHandler) PutTransaction(c echo.Context) error {
	h.logger.Info("handling put transaction request")

	id, err := util.GetParamInt64(c, "transaction_id")
	if err != nil {
		h.logger.Error("invalid transaction id", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid transaction id"})
	}

	var req api.PutTransactionRequest
	if err := c.Bind(&req); err != nil {
		h.logger.Error("failed to decode request body", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	if err := h.service.PutTransaction(id, req); err != nil {
		h.logger.Error("failed to put transaction", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, api.StatusResponse{Status: "ok"})
}

func (h *TransactionHandler) GetTransaction(c echo.Context) error {
	h.logger.Info("handling get transaction request")

	id, err := util.GetParamInt64(c, "transaction_id")
	if err != nil {
		h.logger.Error("invalid transaction id", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid transaction id"})
	}

	tx, err := h.service.GetTransaction(id)
	if err != nil {
		h.logger.Error("failed to get transaction", "error", err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	response := api.GetTransactionResponse{
		Amount:   tx.Amount,
		Type:     tx.Type,
		ParentID: tx.ParentID,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetTransactionsByType(c echo.Context) error {
	h.logger.Info("handling get transactions by type request")

	txType := c.Param("type")
	ids, err := h.service.GetTransactionsByType(txType)
	if err != nil {
		h.logger.Error("failed to get transactions by type", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := api.GetTransactionsByTypeResponse{
		TransactionIds: ids,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetTransactionSum(c echo.Context) error {
	h.logger.Info("handling get transaction sum request")

	id, err := util.GetParamInt64(c, "transaction_id")
	if err != nil {
		h.logger.Error("invalid transaction id", "error", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid transaction id"})
	}

	sum, err := h.service.GetTransactionSum(id)
	if err != nil {
		h.logger.Error("failed to get transaction sum", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := api.GetTransactionSumResponse{
		Sum: sum,
	}

	return c.JSON(http.StatusOK, response)
}
