package usecase

import (
	"github.com/PGabriel20/expenses-go/internal/entity"
)

type GetUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewGetUserUseCase(userRepository entity.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{
		UserRepository: userRepository,
	}
}

func (u *GetUserUseCase) Execute(id string) (UserOutputDto, error) {
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