package main

import "sync"

type TransferRequest struct {
	FromAccountId int     `json:"from_id"`
	ToAccountId   int     `json:"to_id"`
	Amount        float64 `json:"amount"`
}

type Account struct {
	ID      int
	Name    string
	Balance float64
	mu      sync.Mutex // Mutex to ensure atomicity
}

type Bank struct {
	Accounts   map[int]*Account
	mu         sync.Mutex
	generateID func() int
}
