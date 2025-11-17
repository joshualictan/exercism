package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendlist []string, mylist []string) {
	mylist[len(mylist)-1] = friendlist[len(friendlist)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	quantitiesToScale := make([]float64, len(quantities))
	for i, quantity := range quantities {
		quantitiesToScale[i] = quantity * (float64(portions) / 2.0)
	}
	return quantitiesToScale
}
