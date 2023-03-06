package main

import (
	"fmt"
	"os"
)

const (
	buyOption       = "buy"
	fillOption      = "fill"
	takeOption      = "take"
	remainingOption = "remaining"
	exitOption      = "exit"
)

const (
	espressoOption   = "1"
	latteOption      = "2"
	cappuccinoOption = "3"
	backOption       = "back"
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
	beginCoffeeMachineProcess()
}

func beginCoffeeMachineProcess() {
	fmt.Printf("Write action (buy, fill, take, remaining, exit) \n")
	var input string
	fmt.Scan(&input)
	switch true {
	case input == buyOption:
		chooseCoffee()
	case input == fillOption:
		fillMachine()
	case input == takeOption:
		takeMoney()
	case input == remainingOption:
		remainingSupplies()
	case input == exitOption:
		exitProgram()
	default:
		fmt.Printf("I think you got confused... \n")
		beginCoffeeMachineProcess()
	}
}

func chooseCoffee() {
	fmt.Printf("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu: \n")
	var optionChosen string
	fmt.Scan(&optionChosen)
	switch true {
	case optionChosen == espressoOption:
		checkIfEnough(espressoWaterPerCup, espressoMilkPerCup, espressoBeansPerCup)
		updateMachine(-1*espressoWaterPerCup, -1*espressoMilkPerCup, -1*espressoBeansPerCup, -1, espressoCost)
	case optionChosen == latteOption:
		checkIfEnough(latteWaterPerCup, latteMilkPerCup, latteBeansPerCup)
		updateMachine(-1*latteWaterPerCup, -1*latteMilkPerCup, -1*latteBeansPerCup, -1, latteCost)
	case optionChosen == cappuccinoOption:
		checkIfEnough(cappuccinoWaterPerCup, cappuccinoMilkPerCup, cappuccinoBeansPerCup)
		updateMachine(-1*cappuccinoWaterPerCup, -1*cappuccinoMilkPerCup, -1*cappuccinoBeansPerCup, -1, cappuccinoCost)
	case optionChosen == backOption:
		chooseCoffee()
	case optionChosen == remainingOption:
		remainingSupplies()
	default:
		fmt.Printf("I think you got confused... \n")
		beginCoffeeMachineProcess()
	}
	beginCoffeeMachineProcess()
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

	beginCoffeeMachineProcess()
}

func updateMachine(waterUpdate, milkUpdate, beansUpdate, cupsUpdate, moneyUpdate int) {
	water += waterUpdate
	milk += milkUpdate
	beans += beansUpdate
	cups += cupsUpdate
	money += moneyUpdate
}

func takeMoney() {
	fmt.Printf("I gave you $%d \n", money)
	updateMachine(0, 0, 0, 0, -1*money)
	beginCoffeeMachineProcess()
}

func remainingSupplies() {
	fmt.Printf("The coffee machine has: \n")
	fmt.Printf("%d ml of water \n", water)
	fmt.Printf("%d ml of milk \n", milk)
	fmt.Printf("%d g of coffee beans \n", beans)
	fmt.Printf("%d disposable cups \n", cups)
	fmt.Printf("$%d of money \n", money)
	beginCoffeeMachineProcess()
}

func exitProgram() {
	os.Exit(0)
}

func checkIfEnough(waterPerCup, milkPerCup, beansPerCup int) {
	switch true {
	case waterPerCup > water:
		fmt.Printf("Sorry, not enough water! \n")
		beginCoffeeMachineProcess()
	case milkPerCup > milk:
		fmt.Printf("Sorry, not enough milk! \n")
		beginCoffeeMachineProcess()
	case beansPerCup > beans:
		fmt.Printf("Sorry, not enough beans! \n")
		beginCoffeeMachineProcess()
	case cups == 0:
		fmt.Printf("Sorry, not enough cups! \n")
		beginCoffeeMachineProcess()
	default:
		fmt.Printf("I have enough resources, making you a cofee! \n")
	}
}
