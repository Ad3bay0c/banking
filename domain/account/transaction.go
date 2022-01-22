package account

import (
	"github.com/Ad3bay0c/banking/dto"
	"strings"
)

type Transaction struct {
	TransactionID        string `json:"id"`
	AccountID			 string `json:"account_id"`
	Amount				 float64 `json:"amount"`
	TransactionType		 string `json:"transaction_type"`
	TransactionDate      string `json:"transaction_date"`
}

func (transact Transaction) ToNewTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionID:   transact.TransactionID,
		AccountID:       transact.AccountID,
		Amount:          transact.Amount,
		TransactionType: transact.TransactionType,
		TransactionDate: transact.TransactionDate,
	}
}

func (transact *Transaction) IsWithdrawal() bool {
	return strings.ToLower(transact.TransactionType) == dto.WITHDRAW
}
