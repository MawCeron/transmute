package main

var timeUnits = map[string]float64{
	// SI sub-second
	"ns": 1e-9, // nanosecond
	"us": 1e-6, // microsecond
	"ms": 1e-3, // millisecond

	// SI / standard
	"s":   1.0,
	"min": 60.0,
	"h":   3600.0,
	"d":   86400.0,
	"w":   604800.0,

	// Calendar
	"mo":         2629743.83,   // average month (365.25/12 days)
	"y":          31557600.0,   // Julian year (365.25 days)
	"ly_time":    31557600.0,   // same as Julian year (for clarity)
	"decade":     315576000.0,
	"century":    3155760000.0,
	"millennium": 31557600000.0,
	"eon":        3.15576e16,   // 1 billion years

	// Physics
	"planck": 5.391247e-44, // Planck time

	// Practical
	"fortnight": 1209600.0, // 2 weeks
	"shake":     1e-8,      // 10 nanoseconds (nuclear physics slang)
}
