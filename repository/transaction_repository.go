package repository

import (
	"database/sql"
	"log/slog"
	"transaction_service/models"
)

type TransactionRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewTransactionRepository(db *sql.DB, logger *slog.Logger) *TransactionRepository {
	return &TransactionRepository{
		db:     db,
		logger: logger,
	}
}

func (r *TransactionRepository) Put(tx models.Transaction) error {
	r.logger.Info("putting transaction", "id", tx.ID)

	query := `
		INSERT INTO transactions (id, amount, type, parent_id)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE
		SET amount = $2, type = $3, parent_id = $4
	`

	_, err := r.db.Exec(query, tx.ID, tx.Amount, tx.Type, tx.ParentID)
	if err != nil {
		r.logger.Error("failed to put transaction", "error", err)
		return err
	}

	return nil
}

func (r *TransactionRepository) Get(id int64) (*models.Transaction, error) {
	r.logger.Info("getting transaction", "id", id)

	var tx models.Transaction
	err := r.db.QueryRow("SELECT id, amount, type, parent_id FROM transactions WHERE id = $1", id).
		Scan(&tx.ID, &tx.Amount, &tx.Type, &tx.ParentID)
	if err != nil {
		r.logger.Error("failed to get transaction", "error", err)
		return nil, err
	}

	return &tx, nil
}

func (r *TransactionRepository) GetByType(txType string) ([]int64, error) {
	r.logger.Info("getting transactions by type", "type", txType)

	rows, err := r.db.Query("SELECT id FROM transactions WHERE type = $1", txType)
	if err != nil {
		r.logger.Error("failed to get transactions by type", "error", err)
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			r.logger.Error("failed to scan row", "error", err)
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (r *TransactionRepository) GetChildren(parentID int64) ([]int64, error) {
	r.logger.Info("getting child transactions", "parent_id", parentID)

	rows, err := r.db.Query("SELECT id FROM transactions WHERE parent_id = $1", parentID)
	if err != nil {
		r.logger.Error("failed to get child transactions", "error", err)
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			r.logger.Error("failed to scan row", "error", err)
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
