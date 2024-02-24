package usecase

import (
	"errors"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/database/repository"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/model/request"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/model/response"
)

func credit(value int64, id, description string) (*response.ClientTransactionResponse, error) {
	return &response.ClientTransactionResponse{
		Limit:   100000,
		Balance: -9098,
	}, nil
}

func debit(value int64, id, description string) (*response.ClientTransactionResponse, error) {
	transaction, err := repository.DebitBalance(value, id)

	if err != nil {
		return nil, errors.New("failed to debit in the database")
	}

	err = repository.InsertTransaction(id, description, app.DEBIT, value)

	if err != nil {
		return nil, errors.New("failed to insert transaction")
	}

	return response.NewClientTransactionResponse(transaction.Limit, transaction.Balance), nil
}

func TransactionUseCase(id string, req request.ClientTransactionRequest) (*response.ClientTransactionResponse, error) {
	clientExists, err := repository.ClientExists(id)

	if err != nil {
		return nil, errors.New("failed to check if client exists")
	}

	if !clientExists {
		return nil, errors.New("client not found")
	}

	switch req.Type {
	case app.CREDIT:
		return credit(req.Value, id, req.Description)
	case app.DEBIT:
		return debit(req.Value, id, req.Description)
	default:
		// the api will never reach here because req.Type always will be 'c' or 'd'
		return nil, nil
	}
}