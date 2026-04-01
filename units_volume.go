package main

var volumeUnits = map[string]float64{
	// Metric (base: cubic meter)
	"m3":  1,
	"dm3": 1e-3,
	"cm3": 1e-6,
	"mm3": 1e-9,
	"km3": 1e9,

	// Metric capacity
	"l":  1e-3,
	"dl": 1e-4,
	"cl": 1e-5,
	"ml": 1e-6,
	"ul": 1e-9,
	"nl": 1e-12,
	"hl": 0.1,

	// Imperial / US liquid
	"gal":  3.785411784e-3,
	"qt":   9.46352946e-4,
	"pt":   4.73176473e-4,
	"cup":  2.36588236e-4,
	"floz": 2.95735296e-5,
	"tbsp": 1.47867648e-5,
	"tsp":  4.92892159e-6,

	// Imperial / UK
	"gal_uk":  4.54609e-3,
	"qt_uk":   1.1365225e-3,
	"pt_uk":   5.6826125e-4,
	"floz_uk": 2.84130625e-5,
	"gill_uk": 1.420653e-4,

	// US dry
	"bu":     3.523907e-2,
	"pk":     8.809768e-3,
	"dry_qt": 1.101221e-3,
	"dry_pt": 5.506105e-4,

	// Cooking (US)
	"jigger": 4.43603e-5,
	"dash":   6.16115e-7,
	"pinch":  3.08057e-7,

	// Oil & gas
	"bbl":    1.58987295e-1,
	"bbl_uk": 1.63659e-1,
	"mbbl":   158.987295,
	"mmbbl":  158987.295,

	// Astronomical
	"earth_vol": 1.08321e21,
	"sun_vol":   1.41e27,

	// Historical / Traditional
	"arroba_liq": 1.563e-2,
	"fanega":     5.5e-2,
	"almud":      4.625e-3,
	"cuartillo":  1.156e-3,
	"amphora":    2.6e-2,
	"amphora_gr": 3.93e-2,
	"talent_liq": 2.187e-2,
}
