package main

var areaUnits = map[string]float64{
	// Metric
	"m2":  1.0,
	"km2": 1e6,
	"cm2": 1e-4,
	"mm2": 1e-6,
	"um2": 1e-12,
	"ha":  10000.0,  // hectare
	"a":   100.0,    // are
	"da":  10.0,     // deciare

	// Imperial / US
	"ft2":  0.09290304,     // square foot
	"in2":  6.4516e-4,      // square inch
	"yd2":  0.83612736,     // square yard
	"mi2":  2589988.110336, // square mile
	"acre": 4046.8564224,   // acre

	// Nautical
	"nmi2": 3429904.0, // square nautical mile

	// Surveying
	"rood":     1011.7141056, // rood (1/4 acre)
	"township": 93239571.97,  // US survey township (36 square miles)
}
