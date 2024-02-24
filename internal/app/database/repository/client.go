package repository

import (
	"context"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/database"
)

func ClientExists(id string) (bool, error) {
	var clientExists bool
	err := database.GetConn().QueryRow(context.Background(), clientExistsQuery, id).Scan(&clientExists)

	if err != nil {
		return false, err
	}

	return clientExists, nil
}

type TransactionEntity struct {
	Limit   int64
	Balance int64
}

func newTransactionEntity(limit, balance int64) *TransactionEntity {
	return &TransactionEntity{
		Limit:   limit,
		Balance: balance,
	}
}

func DebitBalance(value int64, id string) (*TransactionEntity, error) {
	var limit, balance int64

	err := database.GetConn().QueryRow(context.Background(), debitQuery, value, id, id, value).Scan(&limit, &balance)

	if err != nil {
		return nil, err
	}

	return newTransactionEntity(limit, balance), nil
}

func InsertTransaction(id, description, tipo string, value int64) error {
	_, err := database.GetConn().Exec(context.Background(), insertTransactionQuery, value, tipo, description, id)

	if err != nil {
		return err
	}

	return nil
}
