package main

type Coffee struct {
	water          int
	milk           int
	coffeeBeans    int
	orangeJuice    int
	disposableCups int
	price          int
}

var (
	Espresso   = Coffee{water: 250, coffeeBeans: 16, disposableCups: 1, price: 4}
	Latte      = Coffee{water: 350, milk: 75, coffeeBeans: 20, disposableCups: 1, price: 7}
	Cappuccino = Coffee{water: 200, milk: 100, coffeeBeans: 12, disposableCups: 1, price: 6}
	Bumble     = Coffee{water: 100, orangeJuice: 200, coffeeBeans: 15, disposableCups: 1, price: 7}
)
