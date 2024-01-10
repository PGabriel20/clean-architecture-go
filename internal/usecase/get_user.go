package usecase

import (
	"github.com/PGabriel20/expenses-go/internal/adapter"
)

type GetUserUseCase struct {
	UserRepository adapter.UserRepository
}

func NewGetUserUseCase(userRepository adapter.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		UserRepository: userRepository,
	}
}

func (u *GetUserUseCase) GetUser(id string) (UserOutputDto, error) {
	user, err := u.UserRepository.GetUser(id)

	if err != nil {
		return UserOutputDto{}, err
	}

	outputDto := UserOutputDto{
		ID:         user.ID,
		Username: user.Username,
		Email: user.Email,
	}

	return outputDto, nil
}