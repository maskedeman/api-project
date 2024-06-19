package models

type TransactionLog struct {
	Timestamp

	UserID            uint   `json:"user_id"`            //ID of the user
	TransactionAmount uint   `json:"transaction_amount"` //Amount dealt in current transaction
	RemainingAmount   uint   `json:"remaining_amount"`   //Remaining amount after current transaction
	TransactionType   string `json:"transaction"`        //Type of transaction-deposit, withdraw, transfer
	User              User   `gorm:"foreignKey:user_id"`
}
