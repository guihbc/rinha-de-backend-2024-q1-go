package request

import (
	"errors"
	"strings"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app"
)

type ClientTransactionRequest struct {
	Value       int64  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func (c *ClientTransactionRequest) ValidateFields() error {
	if strings.TrimSpace(c.Type) == "" || (c.Type != app.CREDIT && c.Type != app.DEBIT) {
		return errors.New("o tipo de transacao deve ser c ou d")
	}

	if strings.TrimSpace(c.Description) == "" || len(c.Description) > 10 || len(c.Description) < 1 {
		return errors.New("a descricao deve ter no mínimo 1 caracter e no maximo 10")
	}

	if c.Value <= 0 {
		return errors.New("o valor deve ser um número positivo")
	}

	return nil
}
