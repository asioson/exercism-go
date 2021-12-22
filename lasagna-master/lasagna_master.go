// package lasagna implements methods on how to cook lasagna
package lasagna

// PreparationTime takes in layers and minutes. It returns the estimated total preparation time
func PreparationTime(layers []string, minutes int) int {
    if minutes == 0 {
        minutes = 2
    }
    return len(layers) * minutes
}


// Quantities takes in layers and returns the quantity of noodles and sauce to make lasagna
func Quantities(layers []string) (noodles int, sauce float64) {
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        } else if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return
}

// AddSecretIngredient returns a newList consisting of myList and the last ingridient in friendsList
func AddSecretIngredient(friendsList []string, myList []string) (newList []string) {
    for i := 0; i < len(myList); i++ {
        newList = append(newList, myList[i])
    }
    j := len(friendsList)
    if j > 0 {
        newList = append(newList, friendsList[j-1])
    }
    return
}

// ScaleRecipe takes in amounts and portion and returns neededAmounts for different numbers of portions
func ScaleRecipe(amounts []float64, portions int) (neededAmounts []float64) {
    for i := 0; i < len(amounts); i++ {
        neededAmounts = append(neededAmounts, amounts[i] * float64(portions) * 0.5)
    }
    return
}
