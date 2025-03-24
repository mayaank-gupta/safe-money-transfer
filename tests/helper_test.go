package test

import (
	"testing"

	"github.com/mayaank-gupta/century-pay-assignment/pkg/bank"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test account creation
func TestCreateAccount(t *testing.T) {
	bank := bank.NewBank()

	// Create an account with a balance of 1000
	accountID := bank.CreateAccount("Alex", 2000)

	require.NotZero(t, accountID, "Account ID should not be zero")
	assert.Equal(t, 2000.00, bank.Accounts[accountID].Balance, "Initial balance should be 2000")
}

// Test valid transfer between accounts
func TestValidTransfer(t *testing.T) {
	bank := bank.NewBank()

	acc1 := bank.CreateAccount("Jax", 1000)
	acc2 := bank.CreateAccount("Pearl", 500)

	err := bank.Transfer(acc1, acc2, 200)

	assert.NoError(t, err, "Transfer should succeed")
	assert.Equal(t, 800.00, bank.Accounts[acc1].Balance, "Acc1 should have 800 left")
	assert.Equal(t, 700.00, bank.Accounts[acc2].Balance, "Acc2 should have 700 now")
}

// Test insufficient funds
func TestInsufficientFunds(t *testing.T) {
	bank := bank.NewBank()

	acc1 := bank.CreateAccount("Justin", 100)
	acc2 := bank.CreateAccount("Sindy", 500)

	err := bank.Transfer(acc1, acc2, 200)

	assert.Error(t, err, "Transfer should fail due to insufficient funds")
	assert.Equal(t, 100.00, bank.Accounts[acc1].Balance, "Acc1 balance should remain unchanged")
	assert.Equal(t, 500.00, bank.Accounts[acc2].Balance, "Acc2 balance should remain unchanged")
}

// Test transfer between same account
func TestSelfTransfer(t *testing.T) {
	bank := bank.NewBank()

	acc1 := bank.CreateAccount("Bredy", 1000)

	err := bank.Transfer(acc1, acc1, 200)

	assert.Error(t, err, "Transfer should fail when from and to accounts are the same")
	assert.Equal(t, 1000.00, bank.Accounts[acc1].Balance, "Balance should remain unchanged")
}

// Test transfer to non-existent account
func TestTransferToNonExistentAccount(t *testing.T) {
	bank := bank.NewBank()

	acc1 := bank.CreateAccount("Trex", 1000)
	nonExistentAccount := 99999

	err := bank.Transfer(acc1, nonExistentAccount, 200)

	assert.Error(t, err, "Transfer should fail for non-existent destination account")
	assert.Equal(t, 1000.00, bank.Accounts[acc1].Balance, "Balance should remain unchanged")
}
