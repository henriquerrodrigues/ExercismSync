package meteorology

import "fmt"

type Stringer interface{
    String() string
}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (u TemperatureUnit) String() string{
	units := []string{"°C", "°F"}
    return units[u]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string{
    unit := t.unit.String()
    return fmt.Sprintf("%d %s", t.degree, unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (su SpeedUnit) String() string{
    units := []string{"km/h", "mph"}
    return units[su]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String () string{
	unit:= s.unit.String()
    return fmt.Sprintf("%d %s", s.magnitude, unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string{
	unitT := m.temperature.unit.String()
    unitS := m.windSpeed.unit.String()
    return fmt.Sprintf("%s: %d %s, Wind %s at %d %s, %d%% Humidity", m.location, m.temperature.degree, unitT, m.windDirection, m.windSpeed.magnitude, unitS, m.humidity)
}