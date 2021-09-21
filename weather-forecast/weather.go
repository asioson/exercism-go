// Package weather implements routine that produce 
// formatted message about weather condition in 
// a specific city
package weather

var (
    // CurrentCondition is current weather condition
	CurrentCondition string  
    // CurrentLocation is city name
	CurrentLocation  string  
)

// Forecast takes in arguments city and condition and 
// returns formatted message about the weather condition in specific city
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
