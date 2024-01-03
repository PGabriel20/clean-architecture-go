package adapter

import "github.com/PGabriel20/expenses-go/internal/entity"

type UserRepository interface {
	RegisterUser(user entity.User) error
	GetUser(id string) (*entity.User, error)
	//@TODO - implement
	//AuthenticateUser
	//UpdateUserDetails
	//ChangePassword
	//DeleteUserAccount
}