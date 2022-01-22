package account

import (
	"fmt"
	"github.com/Ad3bay0c/banking/errs"
	"github.com/Ad3bay0c/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

const (
	TABLE = "accounts"
	TRANSACTION = "transactions"
)

type RepositoryDB struct {
	client *sqlx.DB
}

func NewRepositoryDB(DBClient *sqlx.DB) RepositoryDB {
	return RepositoryDB{
		client: DBClient,
	}
}

func (accountDB RepositoryDB) Save(account Account) (*Account, *errs.AppError) {
	query := fmt.Sprintf("INSERT INTO %s(customer_id, opening_date, account_type, amount, status) VALUES($1, $2, $3, $4, $5)", TABLE)
	res, err := accountDB.client.Exec(query, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected error/customer ID does not exists")
	}
	id, _ := res.LastInsertId()
	account.ID = strconv.FormatInt(id, 10)
	return &account, nil
}

func (accountDB RepositoryDB) FindBy(id string) (*Account, *errs.AppError)  {
	var account Account
	query := fmt.Sprintf("SELECT * FROM %s WHERE account_id = $1", TABLE)
	err := accountDB.client.Get(&account, query, id)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &account, nil
}

func (accountDB RepositoryDB) SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)  {
	trx, err := accountDB.client.Begin()
	if err != nil {
		logger.Error("Error while starting database transaction"+err.Error())
        return nil, errs.NewUnexpectedError("unexpected database error")
	}
	query := fmt.Sprintf("INSERT INTO %s(account_id, transaction_type, amount, transaction_date) VALUES($1, $2, $3, $4)", TRANSACTION)
	result, _ := trx.Exec(query, transaction.AccountID, transaction.TransactionType, transaction.Amount, transaction.TransactionDate)


	if transaction.IsWithdrawal() {
		query = fmt.Sprintf("UPDATE %s SET amount = amount - $1 WHERE account_id = $2", TABLE)
		_, err = trx.Exec(query, transaction.Amount, transaction.AccountID)
	} else {
		query = fmt.Sprintf("UPDATE %s SET amount = amount + $1 WHERE account_id = $2", TABLE)
		_, err = trx.Exec(query, transaction.Amount, transaction.AccountID)
	}
	if err != nil {
		logger.Error("Error while saving transaction"+err.Error())
		trx.Rollback()
        return nil, errs.NewUnexpectedError("unexpected database error")
	}

	if err := trx.Commit(); err != nil {
		logger.Error("Error while starting database transaction"+err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	// retrieve the latest account balance
	account, respErr := accountDB.FindBy(transaction.AccountID)
	if respErr != nil {
		return nil, respErr
	}
	transaction.Amount = account.Amount
	id, _ := result.LastInsertId()
	transaction.TransactionID = strconv.FormatInt(id, 10)

	return &transaction, nil
}