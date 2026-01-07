package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(customer *customer, transaction transaction) error {
	var err error
	if transaction.transactionType != transactionDeposit && transaction.transactionType != transactionWithdrawal {
		err = errors.New("unknown transaction type")
		return err
	}
	if transaction.transactionType == transactionDeposit {
		customer.balance += transaction.amount
	} else if transaction.transactionType == transactionWithdrawal && customer.balance < transaction.amount {
		err = errors.New("insufficient funds")
		return err
	} else {
		customer.balance -= transaction.amount
	}
	return nil

}
