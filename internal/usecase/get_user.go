package usecase

import (
	"github.com/PGabriel20/expenses-go/internal/entity"
)

type GetUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewGetUserUseCase(userRepository entity.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		UserRepository: userRepository,
	}
}

func (r *RegisterUserUseCase) ExecuteGet(id string) (UserOutputDto, error) {
	user, err := r.UserRepository.Get(id)

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