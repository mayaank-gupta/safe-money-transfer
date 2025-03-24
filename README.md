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

