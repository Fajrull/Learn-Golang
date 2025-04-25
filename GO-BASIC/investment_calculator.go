package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main(){
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	// fmt.Print("Investment amount: ")
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected return rate: ")
	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// Untuk menampilkan dengan 2 angka di belakang koma menggunakan printf
	// fmt.Printf("Future Value %.2f\n",futureValue)
	// fmt.Println("Future Value (adjusted for inflation) : ", futureRealValue)

	formatedFV := fmt.Sprintf("Future Value : %.1f\n", futureValue)
	formatedRFV := fmt.Sprintf("Future Value (adjusted for inflation) : %.1f\n", futureRealValue)

	fmt.Print(formatedFV, formatedRFV)
}

func outputText(text string){
	fmt.Print(text)
}

// function bisa mengembalikan lebih dari satu nilai
func calculateFutureValues(investmentAmount,expectedReturnRate,years float64) (float64, float64){
	fv := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv := fv / math.Pow(1 + inflationRate / 100, years)
	return fv, rfv
}