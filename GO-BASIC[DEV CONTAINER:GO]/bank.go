package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit : ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount, Must be greater than 0.")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdrawl amount : ")
		var widhdrawlAmount float64
		fmt.Scan(&widhdrawlAmount)

		if widhdrawlAmount <= 0 {
			fmt.Println("Invalid amount, Must be greater than 0.")
			return
		}

		if widhdrawlAmount > accountBalance {
			fmt.Println("Invalid amount, You can't withdraw more than you have.")
			return
		}

		accountBalance -= widhdrawlAmount
		fmt.Println("Balance updated! New amount :", accountBalance)
	} else {
		fmt.Println("GoodBye!")
	}

	fmt.Println("Your choice: ", choice)

}