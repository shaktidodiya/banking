package domain

import (
	"Desktop/golang/banking/errs"
	"Desktop/golang/banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last account id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// start database transaction
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while creating transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error")
	}

	// inserting bank account transaction
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values(?, ?, ?, ?)"
	result, err := tx.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while inserting new transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	// updating account balance
	if t.IsWithdrawal() {
		sqlUpdate := "UPDATE accounts SET amount = amount - ? where account_id = ?"
		_, err = tx.Exec(sqlUpdate, t.Amount, t.AccountId)
	} else {
		sqlUpdate := "UPDATE accounts SET amount = amount + ? where account_id = ?"
		_, err = tx.Exec(sqlUpdate, t.Amount, t.AccountId)
	}

	// in case of error rollback, any changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while committing transaction for bank account " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last transaction id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	// getting the latest account info from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if err != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
