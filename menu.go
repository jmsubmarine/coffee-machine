package main

import "fmt"

func handleMenu(machine *Machine) {
	for {
		action := showMenu()

		switch action {
		case "buy":
			buy(machine)
		case "fill":
			fill(machine)
		case "take":
			take(machine)
		case "remaining":
			printRemaining(machine)
		case "exit":
			return
		default:
			fmt.Println("Unknown action")
		}
	}
}

func showMenu() string {
	var action string
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)
	return action
}
