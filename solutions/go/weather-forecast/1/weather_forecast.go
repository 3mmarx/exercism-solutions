// Package weather provides current forecast information
// for cities in Goblinocus.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the city for which the weather is forecast.
	CurrentLocation string
)

// Forecast returns a description of the current weather condition in the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
