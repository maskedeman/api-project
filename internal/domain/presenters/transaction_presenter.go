package presenters

import "time"

type TransferRequest struct {
	From   uint `json:"from" validate:"required"` //requester id
	Amount uint `json:"amount"`                   //amount to be transferred
	To     uint `json:"to" validate:"required"`   //receiver id
}

type TransferResponse struct {
	TransactionID     uint        `json:"transaction_id"`
	Time              time.Time   `json:"time"`
	FromUser          interface{} `json:"debitedFrom"`
	ToUser            interface{} `json:"creditedTo"`
	TransactionAmount uint        `json:"transaction_amount"`
	RemainingAmount   uint        `json:"remaining_amount"`
	TransactionType   string      `json:"transactionType"`
}

type BalanceResponse struct {
	User           interface{} `json:"user"`
	CurrentBalance uint        `json:"current_balance"`
}
