package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := int(CalculateWorkingCarsPerHour(productionRate, successRate))
	return workingCarsPerHour / 60
}

var costPerTenUnits int = 95000
var costPerSingleUnit int = 10000

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	singleUnits := carsCount % 10
	groupsOfTen := (carsCount - singleUnits) / 10

	return (uint(groupsOfTen) * uint(costPerTenUnits)) + (uint(singleUnits) * uint(costPerSingleUnit))
}
