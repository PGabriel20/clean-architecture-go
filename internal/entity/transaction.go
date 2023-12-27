package entity

import "errors"

var (
	ErrInvalidTransaction         = errors.New("transaction must have an ID")
	ErrInvalidTransactionAccount         = errors.New("transaction must have an account ID")
	ErrInvalidAmmount         = errors.New("amount must not be zero")
	ErrInvalidTransactionType = errors.New("invalid transaction type")
	ErrInvalidTransactionCategory = errors.New("invalid transaction category")
)

// TransactionType represents the type of a transaction (Input/Output).
type TransactionType string

const (
	Input  TransactionType = "Input"
	Output TransactionType = "Output"
)

// TransactionCategory represents the category of a transaction.
type TransactionCategory string

const (
	Groceries    TransactionCategory = "Groceries"
	Utilities    TransactionCategory = "Utilities"
	Entertainment TransactionCategory = "Entertainment"
	Health       TransactionCategory = "Health"
)

type Transaction struct {
	ID       string
	AccountID string
	Amount   float64
	Type     TransactionType
	Category TransactionCategory
}

func NewTransaction(id string, accountID string, amount float64, transactionType TransactionType, category TransactionCategory) (*Transaction, error) {
	transaction := &Transaction{
		ID:       id,
		AccountID: accountID,
		Amount:   amount,
		Type:     transactionType,
		Category: category,
	}

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transaction *Transaction) isValid() error {
	if transaction.ID == "" {
		return ErrInvalidTransaction
	}

	if transaction.AccountID == "" {
		return ErrInvalidTransactionAccount
	}

	if transaction.Amount == 0 {
		return ErrInvalidAmmount
	}

	// Validate TransactionType
	switch transaction.Type {
	case Input, Output:
		// Valid transaction type
	default:
		return ErrInvalidTransactionType
	}

	// Validate TransactionCategory
	switch transaction.Category {
	case Groceries, Utilities, Entertainment, Health:
		// Valid transaction category
	default:
		return ErrInvalidTransactionCategory
	}

	return nil
}
