package main

import (
	"fmt"
	"os"
	"strconv"
)

// Goals
// 1) Validate user input
// 		=> Show error message & exit if invalid input is provided
// 			-no negative numbers
// 			-not 0
// Store calculated results into file

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Revenue :")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate (in percent): ")

	// earningBeforeTax := revenue - expenses
	// profit := earningBeforeTax - (((taxRate) / 100) * earningBeforeTax)
	// ratio := earningBeforeTax / profit
	earningBeforeTax, profit, ratio := calculateEarning(revenue, expenses, taxRate)
	storeResult(earningBeforeTax, profit, ratio)
	// fmt.Println("Earning Before Tax: ", earningBeforeTax)
	// fmt.Println("Profit: ", profit)
	// fmt.Printf("Ratio: %.2f\n", ratio)
	report(earningBeforeTax, profit, ratio)
}

func getUserInput(text string) float64 {
	var userInput string
	fmt.Print(text)
	fmt.Scan(&userInput)
	parsedValue, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		panic("invalid input")
	}

	if parsedValue <= 0 {
		panic("Input can't be zero or negative")
	}
	return parsedValue
}

func calculateEarning(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - ((taxRate / 100) * ebt)
	ratio = ebt / profit

	return
}

func report(ebt, profit, ratio float64) {
	fmt.Printf("Earning Before Tax: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func storeResult(ebt, profit, ratio float64) {
	result := fmt.Sprintf("Earning Before Tax: %v\nProfit: %v\nRatio: %v", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(result), 0644)
}
