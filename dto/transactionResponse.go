package dto

type TransactionResponse struct {
	TransactionID        string `json:"id"`
	AccountID			 string `json:"account_id"`
	Amount				 float64 `json:"amount"`
	TransactionType		 string `json:"transaction_type"`
	TransactionDate      string `json:"transaction_date"`
}
