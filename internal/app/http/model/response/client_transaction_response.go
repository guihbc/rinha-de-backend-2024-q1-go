package response

type ClientTransactionResponse struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}

func NewClientTransactionResponse(limit, balance int64) *ClientTransactionResponse {
	return &ClientTransactionResponse{
		Limit:   limit,
		Balance: balance,
	}
}
