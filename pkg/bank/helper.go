package bank

import (
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mayaank-gupta/century-pay-assignment/pkg/shared"
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

func (b *Bank) CreateAccount(name string, balance float64) int {
	b.mu.Lock()
	defer b.mu.Unlock()
	generatedAccountId := b.generateID()
	b.Accounts[generatedAccountId] = &Account{Name: name, Balance: balance}
	log.Printf("Created account %d with balance %f", generatedAccountId, balance)
	return generatedAccountId
}

func (b *Bank) Transfer(fromId, toId int, amount float64) error {

	if fromId == toId {
		log.Printf("Transfer failed: Source and destination accounts are the same (%d)", fromId)
		return fmt.Errorf("cannot transfer to the same account")
	}

	b.mu.Lock()
	fromAcc, fromExists := b.Accounts[fromId]
	toAcc, toExists := b.Accounts[toId]
	b.mu.Unlock()

	if !fromExists || !toExists {
		log.Printf("Transfer failed: Account %d or %d does not exist", fromId, toId)
		return fmt.Errorf("one or both accounts do not exist")
	}

	fromAcc.mu.Lock()
	defer fromAcc.mu.Unlock()

	if fromAcc.Balance < amount {
		log.Printf("Transfer failed: Insufficient funds in account %d", fromId)
		return fmt.Errorf("insufficient funds")
	}

	toAcc.mu.Lock()
	defer toAcc.mu.Unlock()

	fromAcc.Balance -= amount
	toAcc.Balance += amount
	log.Printf("Transfer successful: %d -> %d (Amount: %.2f)", fromId, toId, amount)
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
	var req shared.TransferRequest
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
