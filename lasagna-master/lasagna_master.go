package lasagna

func PreparationTime(layers []string, average int) int {
	if average == 0 {
		average = 2
	}
	return len(layers) * average
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		}
		if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	newAmounts := make([]float64, len(amounts))
	for i := 0; i < len(amounts); i++ {
		newAmounts[i] = float64(portions) / 2.0 * amounts[i]
	}
	return newAmounts
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
