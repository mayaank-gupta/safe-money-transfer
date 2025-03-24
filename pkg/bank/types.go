package bank

import "sync"

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
