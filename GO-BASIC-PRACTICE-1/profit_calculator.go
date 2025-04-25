package main

import "fmt"

func main( ){
	revenue := getUserInput("Revenue : ")
	expenses := getUserInput("Expenses : ")
	taxRate := getUserInput("Tax Rate : ")

	//EBT : Earning Before Tax (Laba Sebelum Pajak)
	ebt,profit,ratio := calculateFinancials(revenue,expenses,taxRate)

	fmt.Printf("EBT : %.1f\n", ebt)
	fmt.Printf("Profit : %.1f\n", profit)
	fmt.Printf("Ratio : %.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue ,expenses,taxRate float64) (float64,float64,float64){ 
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt,profit,ratio
}