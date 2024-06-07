package main

import (
	"coffeeshop/coffee"
	"fmt"
)

func main() {
	coffees, err := coffee.GetCoffees()
	if err != nil {
		fmt.Println("Error while getting coffeelist ", err)
		return
	}

	fmt.Println("Printing the list of coffees available in the CoffeeShop...")
	for _, element := range coffees.List {
		result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
		fmt.Println(result)
	}
}