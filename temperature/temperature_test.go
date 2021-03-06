// <auto-generated>
// This code was generated by the UnitCodeGenerator tool
//
// Changes to this file will be lost if the code is regenerated
// </auto-generated>

package temperature

import (
	"testing"
	"math"
	"github.com/google/go-cmp/cmp"
)

func withinTolerance() cmp.Option {
	return cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x + y) / 2.0
		return delta / mean < 0.01
	})
}
func TestConvertKnownCelsiusToFahrenheit(t * testing.T) {
    if !cmp.Equal(Celsius.ToFahrenheit(32.0), 89.6, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 89.6, Celsius.ToFahrenheit(32.0));
    }
    if !cmp.Equal(Celsius.ToFahrenheit(0.9), 33.62, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 33.62, Celsius.ToFahrenheit(0.9));
    }
    if !cmp.Equal(Celsius.ToFahrenheit(0.0), 32.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 32.0, Celsius.ToFahrenheit(0.0));
    }
}

func TestConvertKnownCelsiusToKelvin(t * testing.T) {
    if !cmp.Equal(Celsius.ToKelvin(1.0), 274.15, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 274.15, Celsius.ToKelvin(1.0));
    }
    if !cmp.Equal(Celsius.ToKelvin(9.9), 283.05, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 283.05, Celsius.ToKelvin(9.9));
    }
    if !cmp.Equal(Celsius.ToKelvin(100.0), 373.15, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 373.15, Celsius.ToKelvin(100.0));
    }
}

func TestConvertKnownCelsiusToRankine(t * testing.T) {
    if !cmp.Equal(Celsius.ToRankine(900.0), 2111.67, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 2111.67, Celsius.ToRankine(900.0));
    }
    if !cmp.Equal(Celsius.ToRankine(12.0), 513.27, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 513.27, Celsius.ToRankine(12.0));
    }
    if !cmp.Equal(Celsius.ToRankine(-3.0), 486.27, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 486.27, Celsius.ToRankine(-3.0));
    }
}

func TestConvertKnownFahrenheitToCelsius(t * testing.T) {
    if !cmp.Equal(Fahrenheit.ToCelsius(109.0), 42.7778, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 42.7778, Fahrenheit.ToCelsius(109.0));
    }
    if !cmp.Equal(Fahrenheit.ToCelsius(56.9), 13.83333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 13.83333, Fahrenheit.ToCelsius(56.9));
    }
    if !cmp.Equal(Fahrenheit.ToCelsius(102.0), 38.8889, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 38.8889, Fahrenheit.ToCelsius(102.0));
    }
}

func TestConvertKnownFahrenheitToKelvin(t * testing.T) {
    if !cmp.Equal(Fahrenheit.ToKelvin(34.5), 274.5389, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 274.5389, Fahrenheit.ToKelvin(34.5));
    }
    if !cmp.Equal(Fahrenheit.ToKelvin(901.0), 755.928, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 755.928, Fahrenheit.ToKelvin(901.0));
    }
    if !cmp.Equal(Fahrenheit.ToKelvin(23.9), 268.65, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 268.65, Fahrenheit.ToKelvin(23.9));
    }
}

func TestConvertKnownFahrenheitToRankine(t * testing.T) {
    if !cmp.Equal(Fahrenheit.ToRankine(123.0), 582.67, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 582.67, Fahrenheit.ToRankine(123.0));
    }
    if !cmp.Equal(Fahrenheit.ToRankine(9.2), 468.87, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 468.87, Fahrenheit.ToRankine(9.2));
    }
    if !cmp.Equal(Fahrenheit.ToRankine(0.2), 459.87, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 459.87, Fahrenheit.ToRankine(0.2));
    }
}

func TestConvertKnownKelvinToCelsius(t * testing.T) {
    if !cmp.Equal(Kelvin.ToCelsius(80.0), -193.15, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -193.15, Kelvin.ToCelsius(80.0));
    }
    if !cmp.Equal(Kelvin.ToCelsius(0.9), -272.25, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -272.25, Kelvin.ToCelsius(0.9));
    }
    if !cmp.Equal(Kelvin.ToCelsius(305.15), 32.0, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 32.0, Kelvin.ToCelsius(305.15));
    }
}

func TestConvertKnownKelvinToFahrenheit(t * testing.T) {
    if !cmp.Equal(Kelvin.ToFahrenheit(123.4), -237.55, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -237.55, Kelvin.ToFahrenheit(123.4));
    }
    if !cmp.Equal(Kelvin.ToFahrenheit(800.0), 980.33, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 980.33, Kelvin.ToFahrenheit(800.0));
    }
    if !cmp.Equal(Kelvin.ToFahrenheit(99.999), -279.6718, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -279.6718, Kelvin.ToFahrenheit(99.999));
    }
}

func TestConvertKnownKelvinToRankine(t * testing.T) {
    if !cmp.Equal(Kelvin.ToRankine(156.0), 280.8, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 280.8, Kelvin.ToRankine(156.0));
    }
    if !cmp.Equal(Kelvin.ToRankine(8.2), 14.76, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 14.76, Kelvin.ToRankine(8.2));
    }
    if !cmp.Equal(Kelvin.ToRankine(0.8), 1.44, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 1.44, Kelvin.ToRankine(0.8));
    }
}

func TestConvertKnownRankineToCelsius(t * testing.T) {
    if !cmp.Equal(Rankine.ToCelsius(190.0), -167.59444444, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -167.59444444, Rankine.ToCelsius(190.0));
    }
    if !cmp.Equal(Rankine.ToCelsius(0.7), -272.76111111, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -272.76111111, Rankine.ToCelsius(0.7));
    }
    if !cmp.Equal(Rankine.ToCelsius(900.0), 226.85, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 226.85, Rankine.ToCelsius(900.0));
    }
}

func TestConvertKnownRankineToFahrenheit(t * testing.T) {
    if !cmp.Equal(Rankine.ToFahrenheit(109.0), -350.67, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -350.67, Rankine.ToFahrenheit(109.0));
    }
    if !cmp.Equal(Rankine.ToFahrenheit(3567.0), 3107.33, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 3107.33, Rankine.ToFahrenheit(3567.0));
    }
    if !cmp.Equal(Rankine.ToFahrenheit(9.0), -450.67, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", -450.67, Rankine.ToFahrenheit(9.0));
    }
}

func TestConvertKnownRankineToKelvin(t * testing.T) {
    if !cmp.Equal(Rankine.ToKelvin(123.0), 68.333333333, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 68.333333333, Rankine.ToKelvin(123.0));
    }
    if !cmp.Equal(Rankine.ToKelvin(0.9), 0.5, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 0.5, Rankine.ToKelvin(0.9));
    }
    if !cmp.Equal(Rankine.ToKelvin(23.0), 12.777777778, withinTolerance()) {
        t.Fatalf("Expected %f, was %f", 12.777777778, Rankine.ToKelvin(23.0));
    }
}

