package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)

func NewBank() *Bank {
	idGenerator := func() func() int {
		var id = 10000
		var mu sync.Mutex
		return func() int {
			mu.Lock()
			defer mu.Unlock()
			id++
			return id
		}
	}()
	return &Bank{
		Accounts:   make(map[int]*Account),
		generateID: idGenerator,
	}
}

func (b *Bank) CreateAccount(name string, balance float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	generatedAccountId := b.generateID()
	b.Accounts[generatedAccountId] = &Account{Name: name, Balance: balance}
}

func (b *Bank) Transfer(fromId, toId int, amount float64) error {

	b.mu.Lock()
	fromAcc, fromExists := b.Accounts[fromId]
	toAcc, toExists := b.Accounts[toId]
	b.mu.Unlock()

	if !fromExists || !toExists {
		return fmt.Errorf("one or both accounts do not exist")
	}

	fromAcc.mu.Lock()
	defer fromAcc.mu.Unlock()

	if fromAcc.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}

	toAcc.mu.Lock()
	defer toAcc.mu.Unlock()

	fromAcc.Balance -= amount
	toAcc.Balance += amount
	return nil
}

func AccountIdCounter() func() int {
	count := 0
	increment := func() int {
		count++
		return count
	}
	return increment
}

func (b *Bank) TransferHandler(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request"})
		return
	}

	if err := b.Transfer(req.FromAccountId, req.ToAccountId, req.Amount); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Transfer successful"})
}
