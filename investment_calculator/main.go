package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {

	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation: %.2f\n", futureRealValue)

	// outputs information
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation: %.2f\n", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.2f\nFuture Value (adjusted for inflation: %.2f\n`, futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedFRV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	rfv = fv / math.Pow(1+inflationRate/100, float64(years))
	// return fv, rfv
	return
}
