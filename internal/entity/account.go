package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidAccount      = errors.New("an account needs to have an ID")
	ErrInvalidAccountUser  = errors.New("an account needs to have a user")
)

type AccountRepository interface {
	CreateAccount(account Account) (Account, error)
	GetAccount(id string) (Account, error)
	//@TODO - implement
	//CreateAccount
	//CloseAccount
	//GetAccountBalance
}


type Account struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance float64
}

func NewAccount(id uuid.UUID, userId uuid.UUID) (*Account, error) {
	account := &Account{
		ID:     id,
		UserID: userId,
	}

	err := account.IsValid()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (account *Account) IsValid() error {
	if account.ID == uuid.Nil {
		return ErrInvalidAccount
	}

	if account.UserID == uuid.Nil {
		return ErrInvalidAccountUser
	}

	return nil
}

// Domain-level business rules/concerns
func (account *Account) IncreaseBalance(amount float64) error {
	account.Balance += amount
	return nil
}

func (account *Account) DecreaseBalance(amount float64) error {
	account.Balance -= amount
	return nil
}
