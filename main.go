package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	bank := NewBank()

	bank.CreateAccount("Mark", 100.00)
	bank.CreateAccount("Jane", 50.00)
	bank.CreateAccount("Adam", 0.00)

	r := gin.Default()
	r.POST("/transfer", func(c *gin.Context) {
		var req TransferRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		err := bank.Transfer(req.FromAccountId, req.ToAccountId, req.Amount)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
	})
	r.Run(":7070")
}
