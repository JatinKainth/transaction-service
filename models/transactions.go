package models

type Transaction struct {
	ID       int64   `json:"id"`
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"`
	ParentID *int64  `json:"parent_id"`
}
