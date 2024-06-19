package models

type UserBalance struct {
	Timestamp

	UserID uint `json:"user_id"` //ID of the user
	Amount uint `json:"amount"`  //Amount dealt in each transaction
}
