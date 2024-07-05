package entity

type IntentCreationBody struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}
