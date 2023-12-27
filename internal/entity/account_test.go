package entity_test

import (
	"testing"

	"github.com/PGabriel20/expenses-go/internal/entity"
)

//New Account should be created with default balance of 0
func TestAccount_NewAccount(t *testing.T) {
	account, err := entity.NewAccount("123", "John Doe")

	if account.Balance != 0 {
		t.Errorf("expected balance to be 0, but got: %v", account.Balance)
	}

	if err != nil {
		t.Fatalf("expected no errors but got %v", err)
	}
}

//Empty name should return ErrInvalidAccount
func TestAccount_EmptyName(t *testing.T) {
	_, err := entity.NewAccount("123", "")

	if err != entity.ErrInvalidAccount || err == nil {
		t.Errorf("expected error '%v', but got %v", entity.ErrInvalidAccount, err)
	}
}