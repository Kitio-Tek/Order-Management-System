package main

import "context"

type store struct {

	// add here our mongo database
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {

	// add here our mongo create user operation
	return nil
}