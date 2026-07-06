package main

import "fmt"

func main() {
	balance := uint(200)
	getName()
	bet := getBet(balance)
	fmt.Println(bet)
}

func getName() string {
	name := ""

	fmt.Println("Welcome to Slotmachine App")
	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Welcome %s", name)
	return name
}

func getBet(balance uint) uint {
	var bet uint
	for true {
		fmt.Printf("Enter your bet (balance = $%d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("Bet cannot be larger than the balance")
		} else {
			break
		}
	}
	return bet
}
