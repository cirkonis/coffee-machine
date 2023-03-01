package main

import "fmt"

func main() {
	//testies()
	coffee()
}

const (
	waterPerCup = 200
	milkPerCup  = 50
	beansPerCup = 15
)

func coffee() {
	var (
		water int
		milk  int
		beans int
		cups  int
	)
	fmt.Println("Write how many ml of water the coffee machine has: ")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk the coffee machine has: ")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of c the coffee machine has: ")
	fmt.Scan(&beans)
	fmt.Println("Write how many cups of coffee you will need: ")
	fmt.Scan(&cups)
	calculateNeededIngredients(cups)
	minCups := calculateMinCups(water, milk, beans)
	switch true {
	case cups == minCups:
		fmt.Println("Yes, I can make that amount of coffee")
		printCoffeeProcess()
	case cups >= minCups:
		fmt.Printf("No, I can only make %d cups of coffee.\n", minCups)
	case cups < minCups:
		fmt.Printf("Yes, I can make that amount of coffee (and even %d more than that) \n", minCups-cups)
		printCoffeeProcess()
	}
}

func calculateNeededIngredients(cups int) {
	fmt.Printf("For %d cups of coffee you will need: \n", cups)
	fmt.Printf("%d ml of water \n", cups*waterPerCup)
	fmt.Printf("%d ml of milk \n", cups*milkPerCup)
	fmt.Printf("%d g of coffee \n", cups*beansPerCup)
}

func calculateMinCups(water, milk, coffeeBeans int) int {
	minCups := water / waterPerCup
	if milk/milkPerCup < minCups {
		minCups = milk / milkPerCup
	}
	if coffeeBeans/beansPerCup < minCups {
		minCups = coffeeBeans / beansPerCup
	}

	return minCups
}

func printCoffeeProcess() {
	fmt.Println("Starting to make a coffee")
	fmt.Println("Grinding coffee beans")
	fmt.Println("Boiling water")
	fmt.Println("Mixing boiled water with crushed coffee beans")
	fmt.Println("Pouring coffee into the cup")
	fmt.Println("Pouring some milk into the cup")
	fmt.Println("Coffee is ready!")
}

func testies() {
}
