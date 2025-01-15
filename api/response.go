package api

type GetTransactionResponse struct {
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"`
	ParentID *int64  `json:"parent_id"`
}

type GetTransactionsByTypeResponse struct {
	TransactionIds []int64 `json:"transaction_ids"`
}

type GetTransactionSumResponse struct {
	Sum float64 `json:"sum"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
