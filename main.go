package main

import "fmt"

const (
	buyOption  = "buy"
	fillOption = "fill"
	takeOption = "take"
)

const (
	espresso = iota + 1
	latte
	cappuccino
)

const (
	espressoWaterPerCup = 250
	espressoMilkPerCup  = 0
	espressoBeansPerCup = 16
	espressoCost        = 4
)

const (
	latteWaterPerCup = 350
	latteMilkPerCup  = 75
	latteBeansPerCup = 20
	latteCost        = 7
)

const (
	cappuccinoWaterPerCup = 200
	cappuccinoMilkPerCup  = 100
	cappuccinoBeansPerCup = 12
	cappuccinoCost        = 6
)

var (
	water = 400
	milk  = 540
	beans = 120
	cups  = 9
	money = 550
)

func main() {
	printCoffeeStatus()
	handleOption()
}

func printCoffeeStatus() {
	fmt.Printf("The coffee machine has: \n")
	fmt.Printf("%d ml of water \n", water)
	fmt.Printf("%d ml of milk \n", milk)
	fmt.Printf("%d g of coffee beans \n", beans)
	fmt.Printf("%d disposable cups \n", cups)
	fmt.Printf("%d of money \n", money)
}

func handleOption() {
	fmt.Printf("Write action (buy, fill, take)")
	var input string
	fmt.Scan(&input)
	switch true {
	case input == buyOption:
		buyCoffee()
	case input == fillOption:
		fillMachine()
	case input == takeOption:
		takeMoney()
	default:
		fmt.Printf("I think you got confused...")
		handleOption()
	}
}

func buyCoffee() {
	fmt.Printf("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino: \n")
	var drinkChosen int
	fmt.Scan(&drinkChosen)
	switch true {
	case drinkChosen == espresso:
		sellCoffee(espressoWaterPerCup, espressoMilkPerCup, espressoBeansPerCup, espressoCost)
	case drinkChosen == latte:
		sellCoffee(latteWaterPerCup, latteMilkPerCup, latteBeansPerCup, latteCost)
	case drinkChosen == cappuccino:
		sellCoffee(cappuccinoWaterPerCup, cappuccinoMilkPerCup, cappuccinoBeansPerCup, cappuccinoCost)
	}

}

func fillMachine() {
	var (
		waterToAdd int
		milkToAdd  int
		beansToAdd int
		cupsToAdd  int
	)

	fmt.Printf("Write how many ml of water you want to add: \n")
	fmt.Scan(&waterToAdd)
	fmt.Printf("Write how many ml of milk you want to add: \n")
	fmt.Scan(&milkToAdd)
	fmt.Printf("Write how many grams of coffee beans you want to add: \n")
	fmt.Scan(&beansToAdd)
	fmt.Printf("Write how many disposable cups you want to add: \n")
	fmt.Scan(&cupsToAdd)

	updateMachine(waterToAdd, milkToAdd, beansToAdd, cupsToAdd, 0)

	printCoffeeStatus()
	handleOption()
}

func takeMoney() {
	fmt.Printf(" I gave you $%d \n", money)
	updateMachine(0, 0, 0, 0, -1*money)
	printCoffeeStatus()
	handleOption()
}

func sellCoffee(waterPerCup, milkPerCup, beansPerCup, cost int) {
	minCups := calculateMinCups(waterPerCup, milkPerCup, beansPerCup)
	if minCups != 0 {
		updateMachine(-1*waterPerCup, -1*milkPerCup, -1*beansPerCup, -1, cost)
		printCoffeeStatus()
		handleOption()
	} else {
		fmt.Printf("Not enough matierals to make that order. \n")
		fmt.Printf("Try again. \n")
		printCoffeeStatus()
		handleOption()
	}
}

func updateMachine(waterUpdate, milkUpdate, beansUpdate, cupsUpdate, moneyUpdate int) {
	water += waterUpdate
	milk += milkUpdate
	beans += beansUpdate
	cups += cupsUpdate
	money += moneyUpdate
}

func calculateMinCups(waterPerCup, milkPerCup, beansPerCup int) int {
	if milkPerCup == 0 {
		milkPerCup = 1
	}
	minCups := water / waterPerCup
	if milk/milkPerCup < minCups {
		minCups = milk / milkPerCup
	}
	if beans/beansPerCup < minCups {
		minCups = beans / beansPerCup
	}
	if cups < minCups {
		minCups = cups
	}

	return minCups
}
