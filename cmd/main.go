package main

import (
	"go-bankApp/cmd/bank"

	_ "github.com/lib/pq"
)

func main() {
	bank.HandleRequests()
}
