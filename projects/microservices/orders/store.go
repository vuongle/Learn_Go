package main

import "context"

type store struct {
	// add mysql here
}

func NewStore() *store {
	return &store{}
}

//------------------------------------------------
// Implement the OrderStore interface
//------------------------------------------------

//
func (s *store) Create(context.Context) error {
	return nil
}
