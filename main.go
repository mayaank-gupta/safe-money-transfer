package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayaank-gupta/century-pay-assignment/pkg/bank"
	"github.com/mayaank-gupta/century-pay-assignment/pkg/shared"
)

func main() {
	log.Println("Starting the bank server...")
	bank := bank.NewBank()

	bank.CreateAccount("Mark", 100.00)
	bank.CreateAccount("Jane", 50.00)
	bank.CreateAccount("Adam", 0.00)

	log.Println("Test accounts created")

	r := gin.Default()
	r.POST("/transfer", func(c *gin.Context) {
		var req shared.TransferRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("Invalid request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := bank.Transfer(req.FromAccountId, req.ToAccountId, req.Amount)
		if err != nil {
			log.Printf("Transfer error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
		log.Printf("Updated Account Details: %v, %v", bank.Accounts[req.FromAccountId], bank.Accounts[req.ToAccountId])
	})
	log.Println("Server is running on port 7070...")
	r.Run(":7070")
}
