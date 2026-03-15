// Package weather forecasts the current weather condition of various cities in Goblinocu.
package weather

var (
	// CurrentCondition represents the current weather condition of a certain city in Goblinocu.
	CurrentCondition string
	// CurrentLocation represents certain city in Goblinocu.
	CurrentLocation string
)

// Forecast returns a message with the current location and the current weather condition of this.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
