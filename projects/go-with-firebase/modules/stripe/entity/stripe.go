package entity

import "github.com/stripe/stripe-go/v79"

type IntentCreationBody struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type CustomPrice struct {
	ID         string          `json:"id"`
	UnitAmount int64           `json:"unit_amount"`
	Currency   stripe.Currency `json:"currency"`
}

type CustomProduct struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Price       CustomPrice `json:"price"`
}
