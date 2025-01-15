package api

type PutTransactionRequest struct {
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"`
	ParentID *int64  `json:"parent_id"`
}
