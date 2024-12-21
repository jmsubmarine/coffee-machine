package main

import "fmt"

type Machine struct {
	water          int
	milk           int
	coffeeBeans    int
	orangeJuice    int
	disposableCups int
	money          int
}

func newMachine() Machine {
	return Machine{
		water:          400,
		milk:           540,
		coffeeBeans:    120,
		orangeJuice:    300,
		disposableCups: 9,
		money:          550,
	}
}

func printRemaining(machine *Machine) {
	fmt.Printf("\nThe coffee machine has:\n"+
		"%d ml of water\n%d ml of milk\n%d ml of orange juice\n%d g of coffee beans\n"+
		"%d disposable cups\n$%d of money\n\n",
		machine.water, machine.milk, machine.orangeJuice, machine.coffeeBeans, machine.disposableCups, machine.money)
}

func canMakeCoffee(machine *Machine, coffee Coffee) (bool, string) {
	switch {
	case machine.water < coffee.water:
		return false, "water"
	case machine.milk < coffee.milk:
		return false, "milk"
	case machine.orangeJuice < coffee.orangeJuice:
		return false, "orange juice"
	case machine.coffeeBeans < coffee.coffeeBeans:
		return false, "coffee beans"
	case machine.disposableCups < coffee.disposableCups:
		return false, "disposable cups"
	}
	return true, ""
}

func updateMachine(machine *Machine, coffee Coffee) {
	machine.water -= coffee.water
	machine.milk -= coffee.milk
	machine.orangeJuice -= coffee.orangeJuice
	machine.coffeeBeans -= coffee.coffeeBeans
	machine.disposableCups -= coffee.disposableCups
	machine.money += coffee.price
}
