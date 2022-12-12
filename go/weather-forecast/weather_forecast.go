// Package weather, dear Goblins, here be the package.
package weather

// CurrentCondition: A variable for the current weather condition.
var CurrentCondition string

// CurrentLocation: variable specifies the location the CurrentCondition is describing.
var CurrentLocation string

// Forecast: function responsible for writing out the whole forecast e.g. "<CurrentLocation> - current weather condition: <CurrentCondition>".
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
