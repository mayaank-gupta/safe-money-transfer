# 🚀 Century Pay Assignment

A simple concurrent money transfer system built with **Go** and **Gin**.

## 📌 Features
✅ Transfer money between accounts  
✅ Prevents overdrafts (users can't send more than they have)  
✅ Ensures atomic balance updates  
✅ Simple HTTP API  

---

## 🛠️ Installation

### Install Go (if not installed)
Ensure **Go 1.18+** is installed.  
[Download Go](https://go.dev/dl/) and verify the installation:
```sh
go version
```

### Clone the Repository
```sh
git clone https://github.com/mayaank-gupta/century-pay-assignment.git
cd century-pay-assignment
```

### Initialize Go Modules
```sh
go mod tidy
```

### Running the Project
```sh
go run main.go
```

### Run Tests
```sh
go test ./tests -v
```

# 🔒 Concurrency-Safe Data Structures & Locking Strategy

## 1️⃣ Data Structures Used
- **`map[int]*Account`** → A map to store accounts with their IDs as keys.
- **`sync.Mutex`** → Used to ensure thread safety when modifying shared resources.

---

## 2️⃣ Locking Strategy

### 🔒 **Bank-Level Locking**
**Mutex (`b.mu`) at the `Bank` level** is used when:
- Creating new accounts (`CreateAccount`).
- Accessing or modifying the accounts map (`Transfer`).
- Ensuring atomic ID generation (`generateID`).

### 🔒 **Account-Level Locking**
Each **account has its own mutex (`fromAcc.mu`, `toAcc.mu`)** to prevent race conditions when:
- Checking and updating balances.
- Transferring money between accounts.

---

## 🚀 Why This Approach?
✅ **Fine-grained locking**: Instead of locking the entire bank, we lock only the relevant accounts.  
✅ **Prevents deadlocks**: By ensuring consistent lock ordering (`fromAcc` → `toAcc`).  
✅ **Efficient concurrency**: Allows multiple transactions between different accounts to proceed in parallel.  

---

## 🔹 Summary
This implementation ensures that **money transfers are atomic, prevent race conditions, and avoid deadlocks** while maintaining efficient concurrency. 🚀


