package main

import "fmt"

func buy(machine *Machine) {
	fmt.Println("\nWhat do you want to buy? 1 — espresso, 2 — latte, 3 — cappuccino, 4 — bumble, back — to main menu:")
	var coffeeType string
	fmt.Scan(&coffeeType)

	coffees := map[string]Coffee{
		"1": Espresso,
		"2": Latte,
		"3": Cappuccino,
		"4": Bumble,
	}

	if coffeeType == "back" {
		return
	}

	coffee, exists := coffees[coffeeType]
	if !exists {
		fmt.Println("Invalid selection\n")
		return
	}

	if ok, missing := canMakeCoffee(machine, coffee); ok {
		updateMachine(machine, coffee)
		fmt.Println("I have enough resources, making you a coffee!\n")
	} else {
		fmt.Printf("Sorry, not enough %s!\n\n", missing)
	}
}

func fill(machine *Machine) {
	machine.water += getAmount("\nWrite how many ml of water you want to add:")
	machine.milk += getAmount("Write how many ml of milk you want to add:")
	machine.orangeJuice += getAmount("Write how many ml of orange juice you want to add:")
	machine.coffeeBeans += getAmount("Write how many grams of coffee beans you want to add:")
	machine.disposableCups += getAmount("Write how many disposable cups you want to add:")
	fmt.Println()
}

func take(machine *Machine) {
	fmt.Printf("\nI gave you $%d\n\n", machine.money)
	machine.money = 0
}

func getAmount(prompt string) int {
	var amount int
	fmt.Println(prompt)
	fmt.Scan(&amount)
	return amount
}
