package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "Balance.txt"

func ReadFromfile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to read file!")

	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse balance")
	}
	return balance, nil
}
func WriteToFile(balance float64) {
	BalanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(BalanceText), 0644)
}

func main() {
	var CheckBalance, err = ReadFromfile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------------------")
		// panic("Sorry, can't continue the app!")
	}
	fmt.Println("Welcome to the Bank!")
	for { // Here, the for loop is empty and has no statements, this is to make the whole code inside the for loop repeat itself without any conditions and without any number of times until the program is manually stopped.
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var userInput int
		fmt.Print("Your choice : ")
		fmt.Scan(&userInput)

		switch userInput {
		case 1:
			fmt.Println("Your current balance is : ", CheckBalance)
			WriteToFile(CheckBalance)
		case 2:
			var Deposit float64
			fmt.Println("Enter the amount you want to deposit : ")
			fmt.Scan(&Deposit)
			if Deposit <= 0 {
				fmt.Println("Invalid amount, please enter a positive number")
				//return
				continue
			}
			CheckBalance += Deposit // This is like (CheckBalance = CheckBalance + Deposit)
			fmt.Println("Your new balance is : ", CheckBalance, "$")
			WriteToFile(CheckBalance)
		case 3:
			var Withdraw float64
			fmt.Println("Enter the amout you want to withdraw : ")
			fmt.Scan(&Withdraw)
			if Withdraw <= 0 {
				fmt.Println("Invalid amount, please enter a positive number")
				//return
				continue
			}
			if Withdraw > CheckBalance {
				fmt.Println("Invalid amount, you don't have enough money")
				//return
				continue
			}
			CheckBalance -= Withdraw // This is like (CheckBalance = CheckBalance - Withdraw)
			fmt.Println("Your new balance is : ", CheckBalance, "$")
			WriteToFile(CheckBalance)
		case 4:
			fmt.Println("Exiting application...")
			fmt.Println("Thank you for choosing our bank :)")
			return // Here, we are using return to stop the execution of the program rather than using break because break will only stop the execution of the switch statement and not the for loop.
		default:
			fmt.Println("Invalid option, Please choose a valid option between 1 and 4")
			continue
		}
	}
}
