package response

import "time"

type ClientExtractResponse struct {
	Balance          *ExtractBalance       `json:"saldo"`
	LastTransactions []*ExtractTransaction `json:"ultimas_transacoes"`
}

type ExtractBalance struct {
	Total int64     `json:"total"`
	Date  time.Time `json:"data_extrato"`
	Limit int64     `json:"limite"`
}

type ExtractTransaction struct {
	Value       int64     `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	Date        time.Time `json:"realizada_em"`
}
