package repository

import (
	"context"
	"log"
	"time"

	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/database"
	"github.com/guihbc/rinha-de-backend-2024-q1/internal/app/http/model/response"
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

func CreditBalance(value int64, id string) (*TransactionEntity, error) {
	var limit, balance int64

	err := database.GetConn().QueryRow(context.Background(), creditQuery, value, id, id).Scan(&limit, &balance)

	if err != nil {
		return nil, err
	}

	return newTransactionEntity(limit, balance), nil
}

func ClientExtract(id string) (*response.ClientExtractResponse, error) {
	rows, err := database.GetConn().Query(context.Background(), transactionExtractQuery, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	extractResponse := &response.ClientExtractResponse{
		Balance: &response.ExtractBalance{
			Date: time.Now(),
		},
		LastTransactions: []*response.ExtractTransaction{},
	}

	for rows.Next() {
		var extractDate, transactionDate *time.Time
		var total, limit, transactionValue *int64
		var transactionType, transactionDescription *string
		err := rows.Scan(&total, &extractDate, &limit, &transactionValue, &transactionType, &transactionDescription, &transactionDate)

		if err != nil {
			log.Println(err)
			return nil, err
		}

		brazilianZone := time.FixedZone("America/Sao_Paulo", -3*60*60)

		extractResponse.Balance.Total = *total
		extractResponse.Balance.Date = extractDate.In(brazilianZone)
		extractResponse.Balance.Limit = *limit

		var brazilianTransactionDate time.Time
		var t, description string
		var value int64

		var transaction *response.ExtractTransaction

		if transactionDate == nil && transactionDescription == nil && transactionValue == nil && transactionType == nil {
			transaction = nil
		} else {
			brazilianTransactionDate = transactionDate.In(brazilianZone)
			description = *transactionDescription
			value = *transactionValue
			t = *transactionType

			transaction = &response.ExtractTransaction{
				Value:       value,
				Type:        t,
				Description: description,
				Date:        brazilianTransactionDate,
			}

			extractResponse.LastTransactions = append(extractResponse.LastTransactions, transaction)
		}
	}

	err = rows.Err()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return extractResponse, nil
}
