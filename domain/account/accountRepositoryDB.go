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
)

type RepositoryDB struct {
	client *sqlx.DB
}

func NewRepositoryDB(DBClient *sqlx.DB) *RepositoryDB {
	return &RepositoryDB{
		client: DBClient,
	}
}

func (accountDB RepositoryDB) Save(account Account) (*Account, *errs.AppError) {
	query := fmt.Sprintf("INSERT INTO %s(customer_id, opening_date, account_type, amount, status) VALUES($1, $2, $3, $4, $5)", TABLE)
	res, err := accountDB.client.Exec(query, account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected error")
	}
	id, _ := res.LastInsertId()
	account.ID = strconv.FormatInt(id, 10)
	return &account, nil
}
