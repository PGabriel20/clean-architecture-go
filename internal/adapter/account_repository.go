package adapter

import "github.com/PGabriel20/expenses-go/internal/entity"

type AccountRepository interface {
	CreateAccount(account entity.Account) (entity.Account, error)
	GetAccount(id string) (entity.Account, error)
	//@TODO - implement
	//CreateAccount
	//CloseAccount
	//GetAccountBalance
}
