package types

import "fmt"

// Celsius type
type Celsius struct {
	temperature float64
}

// ToFahrenheit converts celsius to Fahrenheit
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit{temperature: (c.temperature * 9 / 5) + 32}
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c.temperature)
}

// Fahrenheit type
type Fahrenheit struct {
	temperature float64
}

// ToCelsius converts Fahrenheit to Celsius
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius{temperature: (f.temperature - 32) * 5 / 9}
}
