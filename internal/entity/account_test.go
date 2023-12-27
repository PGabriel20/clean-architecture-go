package entity_test

import (
	"testing"

	"github.com/PGabriel20/expenses-go/internal/entity"
)

type AccountTestCase struct {
	Test string
	ID string
	UserID string
	Balance float64
	ExpectedErr error
}

//New Account should be created with default balance of 0
func TestAccount_NewAccount(t *testing.T) {
	testCases := []AccountTestCase{
		{
			Test: "Empty account ID",
			ID: "",
			UserID: "123",
			ExpectedErr: entity.ErrInvalidAccount,
		},
		{
			Test: "Empty user ID",
			ID: "123",
			UserID: "",
			ExpectedErr: entity.ErrInvalidAccountUser,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Test, func(t *testing.T) {
			_, err := entity.NewAccount(tc.ID, tc.UserID)

			if err != tc.ExpectedErr {
				t.Errorf("Expected error %v, got %v", tc.ExpectedErr, err)
			}
		})
	}
}

//Empty name should return ErrInvalidAccount
func TestAccount_InitialBalance(t *testing.T) {
	account, _ := entity.NewAccount("123", "1234")

	if account.Balance != 0 {
		t.Errorf("expected balance to be 0, but got: %v", account.Balance)
	}
}