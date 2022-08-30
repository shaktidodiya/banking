package dto

import "Desktop/golang/banking/errs"

type NewAccountRequest struct {
	CustomerId string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open new account you need to deposit at-least 5000.00")
	}
	if r.AccountType != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
