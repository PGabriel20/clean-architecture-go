package usecase

import (
	"github.com/PGabriel20/expenses-go/internal/adapter"
)

type GetUserUseCase struct {
	UserRepository adapter.UserRepository
}

func NewGetUserUseCase(userRepository adapter.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		UserRepository: userRepository,
	}
}

func (r *RegisterUserUseCase) GetUser(id string) (UserOutputDto, error) {
	user, err := r.UserRepository.GetUser(id)

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