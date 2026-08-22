package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int)int{
    if timePerLayer == 0 {
        timePerLayer = 2
    }
    return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    var noodles int
    var sauce float64
    for _, layer := range layers {
        switch layer {
        case "noodles":
            noodles += 50
        case "sauce":
            sauce += 0.2
        }
    }

    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    secretIngredient := friendsList[len(friendsList)-1]
    
    myLastItemIndex := len(myList) - 1
    
    myList[myLastItemIndex] = secretIngredient
    
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64{
    scaleFactor := float64(numPortions) / 2.0
    
    scaledQuantities := make([]float64, len(quantities))
    
    for i, quantity := range quantities {
        scaledQuantities[i] = quantity * scaleFactor
    }
    
    return scaledQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
