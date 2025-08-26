package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return float64(productionRate) * (float64(successRate) / 100.0)
}

// CalculateCarsPerMinute calculates cars produced per minute from produced per hour.
func CalculateCarsPerMinute(productionRate int) float64 {
    return float64(productionRate) / 60.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateCarsPerMinute(productionRate) * (float64(successRate) / 100.0))
}

// CalculateGroups returns the number of full groups of 10 cars and the remaining cars.
func CalculateGroups(carsCount int) (int, int) {
    groups := carsCount / 10
    remaining := carsCount % 10
    return groups, remaining
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    groups, remaining := CalculateGroups(carsCount)
    return uint((groups * 95000) + (remaining * 10000))
}
