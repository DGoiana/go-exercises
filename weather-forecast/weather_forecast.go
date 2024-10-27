// Package weather allow users to check data about the weather.
package weather

// CurrentCondition represents the temperature of a certain location.
var CurrentCondition string

// CurrentLocation represents any location.
var CurrentLocation string

// Forecast returns a string that is a concatenation between it's parameters, a location and the condition of that location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
