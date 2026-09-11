// Package weather represents weather actual.
package weather

var (
    // CurrentCondition represents a certain current condition weather and will be stored.
	CurrentCondition string
    // CurrentLocation represents the position where the condition was metric and will be stored.
	CurrentLocation  string
)
// Forecast represents a message with the current location and how are the condition in this place.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
