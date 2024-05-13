package main

import "context"

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store}
}

//------------------------------------------------
// Implement the OrderService interface
//------------------------------------------------

// Create an order
func (s *service) CreateOrder(context.Context) error {
	return nil
}
