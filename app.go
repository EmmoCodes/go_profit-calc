package main

import (
	"fmt"
)

func main() {
	// const geht auch wie gewohnt -> float64 ist der type
	// var revenue, expenses, taxRate float64

	revenue := getUserInput("Revenue Amount: ")

	expenses := getUserInput("Expenses: ")

	taxRate := getUserInput("Tax Rate: ")

	resultBeforeTax, resultAfterTax, result := calcProfit(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", resultBeforeTax)
	fmt.Printf("%.1f\n", resultAfterTax)
	fmt.Printf("%.1f\n", result)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calcProfit(revenue, expenses, taxRate float64) (float64, float64, float64) {
	resultBeforeTax := revenue - expenses
	resultAfterTax := resultBeforeTax * (1 - taxRate/100)
	result := resultBeforeTax / resultAfterTax

	return resultBeforeTax, resultAfterTax, result
}
