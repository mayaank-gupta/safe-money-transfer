package shared

type TransferRequest struct {
	FromAccountId int     `json:"from_id"`
	ToAccountId   int     `json:"to_id"`
	Amount        float64 `json:"amount"`
}
