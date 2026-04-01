package main

var distanceUnits = map[string]float64{
	// Metric
	"m":   1,
	"dm":  0.1,
	"cm":  0.01,
	"mm":  0.001,
	"dam": 10,
	"hm":  100,
	"km":  1000,

	// Imperial
	"in": 0.0254,
	"ft": 0.3048,
	"yd": 0.9144,
	"mi": 1609.344,

	// Nautical
	"nmi": 1852,

	// Astronomical
	"au":  1.495978707e11,
	"ly":  9.4607304725808e15,
	"pc":  3.085677581e16,
	"kpc": 3.085677581e19,
	"mpc": 3.085677581e22,
	"ls":  2.99792458e8,
	"lm":  1.798754748e10,
	"ld":  2.59020683712e13,

	// Physics
	"pm":  1e-12,
	"fm":  1e-15,
	"ang": 1e-10,
	"nm":  1e-9,
	"um":  1e-6,

	// Historical / Traditional
	"league":          4828.032,
	"league_es":       4179.0,
	"vara_mx":         0.838,
	"vara_es":         0.8359,
	"vara_tx":         0.8467,
	"cubit":           0.4572,
	"cubit_r":         0.5236,
	"fathom":          1.8288,
	"furlong":         201.168,
	"chain":           20.1168,
	"rod":             5.0292,
	"hand":            0.1016,
	"span":            0.2286,
	"pace":            0.762,
	"stadion":         185.0,
	"parasang":        5600.0,
	"link":            0.201168,
	"engineers_chain": 30.48,
}
