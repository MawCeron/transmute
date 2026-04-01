package main

import "fmt"

type tempScale struct {
	toK   func(float64) float64 // converts from this scale to Kelvin
	fromK func(float64) float64 // converts from Kelvin to this scale
}

var temperatureScales = map[string]tempScale{
	"k":  {toK: func(v float64) float64 { return v }, fromK: func(k float64) float64 { return k }},
	"c":  {toK: func(v float64) float64 { return v + 273.15 }, fromK: func(k float64) float64 { return k - 273.15 }},
	"f":  {toK: func(v float64) float64 { return (v + 459.67) * 5 / 9 }, fromK: func(k float64) float64 { return k*9/5 - 459.67 }},
	"r":  {toK: func(v float64) float64 { return v * 5 / 9 }, fromK: func(k float64) float64 { return k * 9 / 5 }},
	"de": {toK: func(v float64) float64 { return 373.15 - v*2/3 }, fromK: func(k float64) float64 { return (373.15 - k) * 3 / 2 }},
	"n":  {toK: func(v float64) float64 { return v*100/33 + 273.15 }, fromK: func(k float64) float64 { return (k - 273.15) * 33 / 100 }},
	"re": {toK: func(v float64) float64 { return v*5/4 + 273.15 }, fromK: func(k float64) float64 { return (k - 273.15) * 4 / 5 }},
	"ro": {toK: func(v float64) float64 { return (v-7.5)*40/21 + 273.15 }, fromK: func(k float64) float64 { return (k-273.15)*21/40 + 7.5 }},
}

// convertTemperature converts a temperature value between two scales.
// Scales: k (Kelvin), c (Celsius), f (Fahrenheit), r (Rankine),
// de (Delisle), n (Newton), re (Reaumur), ro (Romer).
func convertTemperature(value float64, from, to string) (float64, error) {
	src, ok := temperatureScales[from]
	if !ok {
		return 0, fmt.Errorf("unknown unit %q, try --list to see available units", from)
	}
	dst, ok := temperatureScales[to]
	if !ok {
		return 0, fmt.Errorf("unknown unit %q, try --list to see available units", to)
	}
	kelvin := src.toK(value)
	if kelvin < 0 {
		return 0, fmt.Errorf("result is below absolute zero — physically impossible, much like my ex's heart...")
	}
	return dst.fromK(src.toK(value)), nil

}
