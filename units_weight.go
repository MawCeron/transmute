package main

var weightUnits = map[string]float64{
	// Metric
	"kg": 1,
	"g":  1e-3,
	"mg": 1e-6,
	"ug": 1e-9,
	"t":  1000,
	"kt": 1e6,
	"Mt": 1e9,

	// Imperial / US
	"lb":    0.45359237,
	"oz":    0.02834952,
	"tn":    907.18474,
	"tn_uk": 1016.0469,
	"st":    6.35029,
	"gr":    6.47989e-5,
	"dr":    1.77185e-3,

	// Troy (precious metals)
	"oz_t": 0.0311035,
	"lb_t": 0.373242,
	"dwt":  1.55517e-3,

	// Gemstones
	"ct": 2e-4,
	"mp": 2.5e-4,

	// Astronomical
	"m_sun":   1.989e30,
	"m_earth": 5.972e24,
	"m_jup":   1.898e27,

	// Physics
	"u":  1.66054e-27,
	"Da": 1.66054e-27,
	"eV": 1.78266e-36,

	// Historical / Traditional
	"quintal":   100.0,
	"arroba_mx": 11.5,
	"arroba_es": 11.502,
	"libra_es":  0.4601,
	"tonelada":  920.0,
	"talent":    26.0,
	"mina":      0.4333,
	"shekel":    0.01142,
}
