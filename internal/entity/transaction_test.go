package entity_test

import (
	"testing"

	"github.com/PGabriel20/expenses-go/internal/entity"
	"github.com/google/uuid"
)

type TransactionTestCase struct{
	Test string
	ID uuid.UUID
	AccountID uuid.UUID
	Amount   float64
	Type     entity.TransactionType
	Category entity.TransactionCategory
	ExpectedErr error
}

func TestTransaction_NewTransaction(t *testing.T) {
	testCases := []TransactionTestCase{
		{
			Test: "Empty ID",
			ID: uuid.Nil,
			AccountID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			Amount: 100,
			Type: entity.TransactionType(entity.Input),
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidTransaction,
		},
		{
			Test: "Empty account ID",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			AccountID: uuid.Nil,
			Amount: 100,
			Type: entity.TransactionType(entity.Input),
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidTransactionAccount,
		},
		{
			Test: "Zeroed ammount",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			AccountID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			Amount: 0,
			Type: entity.TransactionType(entity.Input),
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidAmmount,
		},
		{
			Test: "Invalid transaction type",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			AccountID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			Amount: 100,
			Type: "InvalidType",
			Category: entity.TransactionCategory(entity.Groceries),
			ExpectedErr: entity.ErrInvalidTransactionType,
		},
		{
			Test: "Invalid transaction category",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			AccountID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
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