package main

var energyUnits = map[string]float64{
	// SI
	"j":  1.0,
	"kj": 1e3,
	"mj": 1e6,
	"gj": 1e9,
	"tj": 1e12,
	"pj": 1e15,

	// Watt-hours
	"wh":  3600.0,
	"kwh": 3.6e6,
	"mwh": 3.6e9,
	"gwh": 3.6e12,
	"twh": 3.6e15,

	// Calories
	"cal":   4.184,    // thermochemical calorie
	"kcal":  4184.0,   // kilocalorie (dietary Calorie)
	"cal15": 4.18580,  // calorie at 15 degrees C

	// Imperial / traditional
	"btu":   1055.05585,  // British thermal unit (ISO)
	"therm": 105480400.0, // US therm (100,000 BTU)
	"ft_lb": 1.35581795,  // foot-pound force

	// Electronics / particle physics
	"ev":  1.602176634e-19,
	"kev": 1.602176634e-16,
	"mev": 1.602176634e-13,
	"gev": 1.602176634e-10,
	"tev": 1.602176634e-7,

	// CGS
	"erg": 1e-7,
}
