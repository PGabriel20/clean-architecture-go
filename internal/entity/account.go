package entity

import "errors"

var (
	ErrInvalidAccount = errors.New("an account needs to have a name")
)

type Account struct {
	ID string
	Name string
	Balance float64
}

func NewAccount(id string, name string) (*Account, error) {
	account := &Account{
		ID: id,
		Name: name,
	}

	err := account.IsValid()

	if err != nil {
		return nil, err
	}

	return account, nil
}

//Fine business rules
func (account *Account) IsValid() error {
	//account shoud have a name
	if account.Name == "" {
		return ErrInvalidAccount
	}

	return nil
}