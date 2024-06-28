package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/schtvr/crypto-converter/internal/handlers"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <USD_AMOUNT> <CRYPTO1> <CRYPTO2>", os.Args[0])
	}

	usdAmount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Invalid USD amount: %v", err)
	}

	crypto1 := os.Args[2]
	crypto2 := os.Args[3]

	result, err := handlers.ConvertHoldings(usdAmount, crypto1, crypto2)
	if err != nil {
		log.Fatalf("Error converting holdings: %v", err)
	}

	fmt.Println(result)
}
