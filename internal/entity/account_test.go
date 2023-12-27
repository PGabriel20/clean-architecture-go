package entity_test

import (
	"testing"

	"github.com/PGabriel20/expenses-go/internal/entity"
	"github.com/google/uuid"
)

type AccountTestCase struct {
	Test string
	ID uuid.UUID
	UserID uuid.UUID
	Balance float64
	ExpectedErr error
}

//New Account should be created with default balance of 0
func TestAccount_NewAccount(t *testing.T) {
	testCases := []AccountTestCase{
		{
			Test: "Empty account ID",
			ID: uuid.Nil,
			UserID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			ExpectedErr: entity.ErrInvalidAccount,
		},
		{
			Test: "Empty user ID",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			UserID: uuid.Nil,
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
	uuid := uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")
	account, _ := entity.NewAccount(uuid, uuid)

	if account.Balance != 0 {
		t.Errorf("expected balance to be 0, but got: %v", account.Balance)
	}
}