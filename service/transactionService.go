package service

import (
	"github.com/Ad3bay0c/banking/domain/account"
	"github.com/Ad3bay0c/banking/dto"
	"github.com/Ad3bay0c/banking/errs"
	"time"
)

func (d DefaultAccountService) SaveTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	if transactionErr := request.Validate(); transactionErr != nil {
		return nil, transactionErr
	}

	// Checks if transaction type is withdrawal
	if request.IsTransactionWithdrawal() {
		acc, err := d.repo.FindBy(request.AccountID) // get the account details of the accountID
		if err != nil {
			return nil, err
		}
		// checks if the account has enough balance to withdraw
		if !acc.CanWithdraw(request.Amount) {
			return nil, errs.NewValidationError("insufficient balance")
		}
	}
	// if all is well, then create a transaction model that will be sent to database access method
	newTransaction := account.Transaction{
		TransactionID:   "",
		AccountID:       request.AccountID,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	// save transaction to database and get a response or error
	response, err := d.repo.SaveTransaction(newTransaction)
	if err != nil {
		return nil, err
	}
	// changed the response to dto.TransactionResponse
	resp := response.ToNewTransactionResponseDto()
	return &resp, nil
}

