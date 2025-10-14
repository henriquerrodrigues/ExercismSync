package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    ratePerHour := float64(productionRate)
    propSuccess := successRate / 100
	var calc float64 = ratePerHour * propSuccess
    return calc
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	ratePerMin := float64(productionRate)
    ratePerMin = ratePerMin/60
    propSuccess := successRate/100
    
	var calcMin float64 = ratePerMin * propSuccess
    return int(calcMin)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    const costOfTen uint = 95000
    const costOfOne uint = 10000
    var dec uint = uint(carsCount/10)
	var rest uint = uint(carsCount%10)

    var total uint = (dec * costOfTen) + (rest * costOfOne)

    return total 
}
