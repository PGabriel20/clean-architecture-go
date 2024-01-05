package usecase

import (
	"github.com/PGabriel20/expenses-go/internal/entity"
	"github.com/PGabriel20/expenses-go/internal/pkg"
	"github.com/google/uuid"
)

type UserInputDto struct {
	Username string
	Email string
	Password string
}

type UserOutputDto struct {
	ID uuid.UUID
	Username string
	Email string
}

type RegisterUserUseCase struct {
	UserRepository entity.UserRepository
}

func NewRegisterUserUseCase(userRepository entity.UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		UserRepository: userRepository,
	}
}

func (r *RegisterUserUseCase) Execute(input UserInputDto) (UserOutputDto, error) {
	hashedPassword, err := pkg.HashPassword(input.Password)

	if err != nil {
		return UserOutputDto{}, err
	}

	user, err := entity.NewUser(uuid.New(), input.Username, input.Email, hashedPassword)

	if err != nil {
		return UserOutputDto{}, err
	}

	if err := r.UserRepository.RegisterUser(*user); err != nil {
		return UserOutputDto{}, err
	}

	outputDto := UserOutputDto{
		ID:         user.ID,
		Username: user.Username,
		Email: user.Email,
	}

	return outputDto, nil
}