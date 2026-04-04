package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    // return the successRate percent of productionRate
    numberOfSuccesfullCars := (float64(productionRate) / 100) * successRate
    return numberOfSuccesfullCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    numberOfSuccesfullCars := (float64(productionRate) / 100) * successRate
    return int(numberOfSuccesfullCars) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupOfTen := carsCount / 10
	unit := carsCount % 10
    return uint((groupOfTen * 95000) + (unit * 10000)) 
}
