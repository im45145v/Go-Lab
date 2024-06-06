package coffee

var coffees = map[string]float32{"Latte": 2.5, "Cappuccino": 2.75, "Flat White": 2.25}

func GetCoffeess() map[string]float32 {
	return coffees
}
