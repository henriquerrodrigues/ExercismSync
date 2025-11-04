package thefarm

import "fmt"

type InvalidCowsError struct{
    message string
    details string
}

func (i *InvalidCowsError) Error() string{
    return fmt.Sprintf("%s: %s", i.message, i.details)
}

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, cows int) (float64, error){
    totalFodder, err := calculator.FodderAmount(cows)

    if err != nil{
        return 0, err
    }

    fatteningFactor, err := calculator.FatteningFactor()

    if err != nil{
        return 0, err
    }

    foodPerCow := (totalFodder * fatteningFactor) / float64(cows)

    return foodPerCow, nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calculator FodderCalculator,cows int) (float64, error){
    if cows > 0{
        return DivideFood(calculator, cows)
    }else{
        return 0, fmt.Errorf("invalid number of cows")
    }
}


// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error{
	if cows < 0{
        return &InvalidCowsError{
        message: fmt.Sprintf("%d cows are invalid", cows),
        details:"there are no negative cows",
    	}
    }
    if cows == 0{
        return &InvalidCowsError{
        message: fmt.Sprintf("%d cows are invalid", cows),
        details:"no cows don't need food",
    	}
    }
    
    return nil
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
