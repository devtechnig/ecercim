// Package weather forecast the currentweather of different cities in Goblinocus.
package weather

var (
    // CurrentCondition holds the name of the current condition (rainy, sunny, winter, etc.).
	CurrentCondition string
    // CurrentLocation holds the name of the current city.
	CurrentLocation  string
)

// Forecast takes two argurments city and condition and return the current condition and the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
