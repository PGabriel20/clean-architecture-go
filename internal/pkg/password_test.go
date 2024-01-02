package pkg_test

import (
	"regexp"
	"testing"

	"github.com/PGabriel20/expenses-go/internal/pkg"
)

//Password should have valid hash
func TestPassword_Hashing(t *testing.T) {
	password := "securePassword"

	hash, err := pkg.HashPassword(password)

	if err != nil {
		t.Fatalf("Error while hashing password: %v", err)
	}

	if len(hash) == 0 {
		t.Error("Hashed password is empty")
	}

	bcryptRegex := regexp.MustCompile(`^\$2[ayb]\$[0-9]{1,2}\$[./A-Za-z0-9]{53}$`)
	isValidHash := bcryptRegex.MatchString(hash)

	if !isValidHash {
		t.Errorf("Password wasnt hashed correctly")
	}
}