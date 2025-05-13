package main

import "context"

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service { //return pointer to service
	return &service{store} //a new instance, pass in store
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}  