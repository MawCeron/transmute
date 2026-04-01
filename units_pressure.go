package main

var pressureUnits = map[string]float64{
	// SI
	"pa":  1.0,
	"hpa": 100.0,
	"kpa": 1000.0,
	"mpa": 1e6,
	"gpa": 1e9,

	// Bar
	"bar":  100000.0,
	"mbar": 100.0,
	"ubar": 0.1,

	// Atmosphere
	"atm": 101325.0,

	// Pounds per square inch
	"psi": 6894.757,
	"ksi": 6894757.0,

	// Mercury / water columns
	"torr":  133.322,
	"mmhg":  133.322,  // same as torr
	"inhg":  3386.389,
	"cmhg":  1333.22,
	"inh2o": 249.089,
	"cmh2o": 98.0665,
	"mmh2o": 9.80665,

	// Technical atmosphere
	"at": 98066.5, // kgf/cm2

	// Historical
	"barye": 0.1, // CGS unit of pressure
}
