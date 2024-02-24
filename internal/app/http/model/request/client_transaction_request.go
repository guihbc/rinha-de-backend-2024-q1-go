package request

import (
	"errors"
	"strings"
)

type ClientTransactionRequest struct {
	Value       int64  `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

func (c *ClientTransactionRequest) ValidateFields() error {
	if strings.TrimSpace(c.Type) == "" || (c.Type != "c" && c.Type != "d") {
		return errors.New("o tipo de transacao deve ser c ou d")
	}

	if strings.TrimSpace(c.Type) == "" || len(c.Description) > 10 || len(c.Description) < 1 {
		return errors.New("a descricao deve ter no mínimo 1 caracter e no maximo 10")
	}

	if c.Value <= 0 {
		return errors.New("o valor deve ser um número positivo")
	}

	return nil
}
