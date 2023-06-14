package lasagna

// PreparationTime() function
func PreparationTime(layers []string, minutes int) (prepTime int) {
	if minutes == 0 {
		minutes = 2
	}
	prepTime = len(layers) * minutes
	return
}

// Quantities() function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}

	return
}

// AddSecretIngredient() function
func AddSecretIngredient(friendsList []string, myList []string) (myNewList []string) {

	myList[len(myList)-1] = friendsList[len(friendsList)-1]
	return
}

// ScaleRecipe() function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaleBy := float64(portions) / 2
	newQuantities := make([]float64, len(quantities))
	for index, amount := range quantities {
		newQuantities[index] = amount * scaleBy
	}
	return newQuantities
}
