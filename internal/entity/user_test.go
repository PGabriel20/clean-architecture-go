package entity_test

import (
	"testing"

	"github.com/PGabriel20/expenses-go/internal/entity"
	"github.com/google/uuid"
)

type UserTestCase struct {
	Test string
	ID uuid.UUID
	Username string
	Email string
	Password string
	ExpectedErr error
}


//User validation errors
func TestUser_NewUser(t *testing.T) {
	testCases := []UserTestCase{
		{
			Test: "Empty ID",
			ID: uuid.Nil,
			Username: "John doe",
			Email: "john.doe@mail.com",
			Password: "123",
			ExpectedErr: entity.ErrInvalidID,
		},
		{
			Test: "Empty username",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			Username: "",
			Email: "john.doe@mail.com",
			Password: "123",
			ExpectedErr: entity.ErrInvalidUsername,
		},
		{
			Test: "Empty email",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			Username: "John doe",
			Email: "",
			Password: "123",
			ExpectedErr: entity.ErrInvalidEmail,
		},
		{
			Test: "Empty password",
			ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			Username: "John Doe",
			Email: "john.doe@mail.com",
			Password: "",
			ExpectedErr: entity.ErrInvalidPassword,
		},
	}


	for _, tc := range testCases {
		t.Run(tc.Test, func(t *testing.T) {
			_, err := entity.NewUser(tc.ID, tc.Username, tc.Email, tc.Password)

			if err != tc.ExpectedErr {
				t.Errorf("Expected error %v, got %v", tc.ExpectedErr, err)
			}
		})
	}
}