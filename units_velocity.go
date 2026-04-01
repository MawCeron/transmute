package main

var velocityUnits = map[string]float64{
	// SI
	"mps": 1.0,       // meters per second
	"kmh": 1.0 / 3.6, // kilometers per hour
	"kph": 1.0 / 3.6, // alias for kmh

	// Imperial
	"mph": 0.44704,  // miles per hour
	"fps": 0.3048,   // feet per second
	"fpm": 0.00508,  // feet per minute
	"ips": 0.0254,   // inches per second

	// Nautical
	"kn": 0.514444, // knots

	// Physics / reference
	"mach": 340.29,      // speed of sound at sea level, 15 degrees C
	"c":    299792458.0, // speed of light in vacuum

	// Astronomical
	"kms": 1000.0, // km/s (used for stellar velocities)
}
