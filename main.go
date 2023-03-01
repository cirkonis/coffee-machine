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

func testies() {
	var number int
	fmt.Scanf("%d", &number)
	switch true {
	case number < 0:
		fmt.Print("Negative!")
	case number == 0:
		fmt.Print("Zero!")
	case number > 0:
		fmt.Print("Positive!")
	}
}

func coffee() {
	gatherInfo()
	printCoffeeProccess()
}

func gatherInfo() {
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
	calcNeeded(cups)
	calculateMinCups(water, milk, beans, cups)
}

func calcNeeded(cups int) {
	fmt.Printf("For %d cups of coffee you will need: \n", cups)
	fmt.Printf("%d ml of water \n", cups*waterPerCup)
	fmt.Printf("%d ml of milk \n", cups*milkPerCup)
	fmt.Printf("%d g of coffee \n", cups*beansPerCup)
}

func printCoffeeProccess() {
	fmt.Println("Starting to make a coffee")
	fmt.Println("Grinding coffee beans")
	fmt.Println("Boiling water")
	fmt.Println("Mixing boiled water with crushed coffee beans")
	fmt.Println("Pouring coffee into the cup")
	fmt.Println("Pouring some milk into the cup")
	fmt.Println("Coffee is ready!")
}

func calculateMinCups(water int, milk int, beans int, cups int) {
	var (
		minCupsWater = (water - waterPerCup*cups) / waterPerCup
		minCupsMilk  = (milk - milkPerCup*cups) / milkPerCup
		minCupsBeans = (beans - beansPerCup*cups) / beansPerCup
		minCups      int
	)

	// Check the lowest number
	if minCupsWater <= minCupsMilk && minCupsWater <= minCupsBeans {
		minCups = minCupsWater
	} else if minCupsMilk <= minCupsWater && minCupsMilk <= minCupsBeans {
		minCups = minCupsMilk
	} else {
		minCups = minCupsBeans
	}

	switch true {
	case minCups == cups:
		fmt.Println("Yes, I can make that amount of coffee")
	case minCups < 0 && cups+minCups > 0:
		fmt.Printf("No, I can make only %d cups \n", minCups+cups)
	case minCups > cups:
		fmt.Printf("Yes, I can make that amount of coffe (and even %d more than that)", minCups-cups)
	case minCups < 0 && cups+minCups < 0:
		fmt.Println("No, I can only make 0 cups of coffee")
	}

}
