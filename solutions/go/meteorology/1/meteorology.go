// Package meteorology adds suitable String methods to all types.
package meteorology

import "strconv"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
// String returns temperature unit as a string.
func (m TemperatureUnit) String() string {
	unit := []string{"°C", "°F"}
	return unit[m]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
// String returns temperature value and unit as a string.
func (t Temperature) String() string {
	return strconv.Itoa(t.degree) + " " + t.unit.String()
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
// String returns speed unit as a string.
func (s SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
// String returns speed value and unit as a string.
func (s Speed) String() string {
	return strconv.Itoa(s.magnitude) + " " + s.unit.String()
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
// String resturns a friendly string that represents MeteorologyData.
func (m MeteorologyData) String() string {
	return m.location + ": " + m.temperature.String() + ", Wind " + m.windDirection + " at " +
		m.windSpeed.String() + ", " + strconv.Itoa(m.humidity) + "% Humidity"
}
