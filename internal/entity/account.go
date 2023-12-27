package entity

import "errors"

var (
	ErrInvalidAccount = errors.New("an account needs to have an ID")
	ErrInvalidAccountUser = errors.New("an account needs to have a user")
)

type Account struct {
	ID string
	UserID string
	Balance float64
}

func NewAccount(id string, userId string) (*Account, error) {
	account := &Account{
		ID: id,
		UserID: userId,
	}

	err := account.IsValid()

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (account *Account) IsValid() error {
	//account shoud be related to a user
	if account.ID == "" {
		return ErrInvalidAccount
	}

	if account.UserID == "" {
		return ErrInvalidAccountUser
	}

	return nil
}

//Domain-level business rules/concerns
func (account *Account) IncreaseBalance(amount float64) error {
	account.Balance += amount
	return nil
}

func (account *Account) DecreaseBalance(amount float64) error {
	account.Balance -= amount
	return nil
}
