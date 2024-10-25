package main

import (
	"fmt"
)

func main(){
var revenue,expenses,taxRate float64

fmt.Print("Revenue Amount: ")
fmt.Scan(&revenue)

fmt.Print("Expenses Amount: ")
fmt.Scan(&expenses)

fmt.Print("Tax Rate: ")
fmt.Scan(&taxRate)

resultBeforeTax := revenue - expenses
resultAfterTax := resultBeforeTax * (1 - taxRate/100)
result := resultBeforeTax / resultAfterTax

// fmt.Println("Profit before tax: ", resultBeforeTax)
// fmt.Println("Profit after tax: ", resultAfterTax)
fmt.Printf("Profit before tax: %.1f\n Profit after tax: %.1f\n Result: %.1f",resultBeforeTax,resultAfterTax,result)

}