// Package weather provides tools to describe the weather forecast for a given location.
package weather

// CurrentCondition describes the current weather condition (e.g., "sunny", "rainy").
var CurrentCondition string

// CurrentLocation indicates the location for which the weather is being reported.
var CurrentLocation string

// Forecast returns a formatted weather report for the given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
