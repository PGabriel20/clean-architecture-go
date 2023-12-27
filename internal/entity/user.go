package entity

import "errors"

type User struct {
	ID string
	Username string
	Email string
	Password string
}

var (
	ErrInvalidID = errors.New("a user must have an ID")
	ErrInvalidUsername = errors.New("a user must have a valid username")
	ErrInvalidEmail = errors.New("a user must have a valid email")
	ErrInvalidPassword = errors.New("a user must have a valid password")
)

func NewUser(id string, username string, email string, password string) (*User, error) {
	user := &User{
		ID: id,
		Username: username,
		Email: email,
		Password: password,
	}

	err := user.isValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}

//Email validation and password hashing are applications concerns, not domains
func (u *User) isValid() error {
		if u.ID == "" {
			return ErrInvalidID
		}
 
		if u.Username == "" {
			return ErrInvalidUsername
		}

		if u.Email == "" {
			return ErrInvalidEmail
		}
		
		if u.Password == "" {
			return ErrInvalidPassword
		}

	return nil
}