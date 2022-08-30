package dto

import "Desktop/golang/banking/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t TransactionRequest) IsTransactionTypeDeposit() bool {
	return t.TransactionType == DEPOSIT
}

func (t TransactionRequest) Validate() *errs.AppError {
	if !t.IsTransactionTypeWithdrawal() && !t.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type should be withdrawal or deposit")
	}

	if t.Amount < 0 {
		return errs.NewValidationError("Amount should be greater than zero")
	}
	return nil
}
