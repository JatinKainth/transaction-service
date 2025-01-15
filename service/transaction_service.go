package service

import (
	"fmt"
	"log/slog"
	"transaction_service/api"
	"transaction_service/models"
	"transaction_service/repository"
)

type TransactionService struct {
	repo   *repository.TransactionRepository
	logger *slog.Logger
}

func NewTransactionService(repo *repository.TransactionRepository, logger *slog.Logger) *TransactionService {
	return &TransactionService{
		repo:   repo,
		logger: logger,
	}
}

func (s *TransactionService) PutTransaction(id int64, req api.PutTransactionRequest) error {
	s.logger.Info("putting transaction", "id", id)

	if req.ParentID != nil {
		_, err := s.repo.Get(id)
		if err == nil {
			if err := s.checkForCycles(id, *req.ParentID); err != nil {
				s.logger.Error("cycle detected in transaction relationships",
					"transaction_id", id,
					"parent_id", *req.ParentID,
					"error", err,
				)
				return fmt.Errorf("invalid parent relationship: %w", err)
			}
		}
	}

	tx := models.Transaction{
		ID:       id,
		Amount:   req.Amount,
		Type:     req.Type,
		ParentID: req.ParentID,
	}
	return s.repo.Put(tx)
}

func (s *TransactionService) checkForCycles(transactionID, parentID int64) error {
	if transactionID == parentID {
		return fmt.Errorf("transaction cannot be its own parent")
	}

	visited := make(map[int64]bool)

	// Start from the potential parent and traverse up
	// If we find the transaction ID, it means there would be a cycle
	current := parentID
	for current != 0 {
		if visited[current] {
			return fmt.Errorf("cycle detected in transaction relationships")
		}
		visited[current] = true

		if current == transactionID {
			return fmt.Errorf("adding this parent would create a cycle")
		}

		tx, err := s.repo.Get(current)
		if err != nil {
			return fmt.Errorf("failed to get transaction %d: %w", current, err)
		}

		if tx.ParentID == nil {
			break
		}
		current = *tx.ParentID
	}

	return nil
}

func (s *TransactionService) GetTransaction(id int64) (*models.Transaction, error) {
	s.logger.Info("getting transaction", "id", id)
	return s.repo.Get(id)
}

func (s *TransactionService) GetTransactionsByType(txType string) ([]int64, error) {
	s.logger.Info("getting transactions by type", "type", txType)
	return s.repo.GetByType(txType)
}

func (s *TransactionService) GetTransactionSum(id int64) (float64, error) {
	s.logger.Info("calculating transaction sum", "id", id)

	visited := make(map[int64]bool)
	return s.calculateSum(id, visited)
}

func (s *TransactionService) calculateSum(id int64, visited map[int64]bool) (float64, error) {
	if visited[id] {
		return 0, nil
	}
	visited[id] = true

	tx, err := s.repo.Get(id)
	if err != nil {
		return 0, err
	}

	sum := tx.Amount

	children, err := s.repo.GetChildren(id)
	if err != nil {
		return 0, err
	}

	for _, childID := range children {
		childSum, err := s.calculateSum(childID, visited)
		if err != nil {
			return 0, err
		}
		sum += childSum
	}

	return sum, nil
}
