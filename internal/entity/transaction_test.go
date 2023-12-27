package entity_test

import (
	"testing"

	"github.com/PGabriel20/expenses-go/internal/entity"
)

type TransactionTestCase struct{
	Test string
	ID string
	AccountID string
	Amount   float64
	Type     entity.TransactionType
	Category entity.TransactionCategory
	ExpectedErr error
}

func TestTransaction_NewTransaction(t *testing.T) {
	testCases := []TransactionTestCase{
		{
			Test: "Empty ID",
			ID: "",
			AccountID: "123",
			Amount: 100,
			Type: entity.TransactionType(entity.Input),
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidTransaction,
		},
		{
			Test: "Empty user ID",
			ID: "123",
			AccountID: "",
			Amount: 100,
			Type: entity.TransactionType(entity.Input),
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidTransactionAccount,
		},
		{
			Test: "Zeroed ammount",
			ID: "123",
			AccountID: "12345",
			Amount: 0,
			Type: entity.TransactionType(entity.Input),
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidAmmount,
		},
		{
			Test: "Invalid transaction type",
			ID: "123",
			AccountID: "12345",
			Amount: 100,
			Type: "InvalidType",
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidTransactionType,
		},
		{
			Test: "Invalid transaction category",
			ID: "123",
			AccountID: "12345",
			Amount: 100,
			Type: entity.TransactionType(entity.Input),
			Category: "InvalidCategory",
			ExpectedErr: entity.ErrInvalidTransactionCategory,
		},	
	}

	for _, tc := range testCases {
		t.Run(tc.Test, func(t *testing.T) {
			_, err := entity.NewTransaction(tc.ID, tc.AccountID, tc.Amount, tc.Type, tc.Category)

			if err != tc.ExpectedErr {
				t.Errorf("Expected error %v, got %v", tc.ExpectedErr, err)
			}
		})
	}
}