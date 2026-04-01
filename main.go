package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var unitCategories = map[string]map[string]float64{
	"area":        areaUnits,
	"data":        dataUnits,
	"distance":    distanceUnits,
	"energy":      energyUnits,
	"pressure":    pressureUnits,
	"time":        timeUnits,
	"velocity":    velocityUnits,
	"volume":      volumeUnits,
	"weight":      weightUnits,
	"temperature": nil, // special case: additive offsets, handled by convertTemperature
}

func main() {
	var (
		category  string
		verbose   bool
		list      bool
		precision int
	)
	flag.StringVar(&category, "c", "distance", "unit category")
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.BoolVar(&list, "list", false, "list available units for the selected category")
	flag.IntVar(&precision, "p", 4, "number of significant digits (default 4)")

	flag.Usage = func() {
		// Scan os.Args for -c value to show category-specific help.
		cat := ""
		for i, arg := range os.Args[1:] {
			if arg == "-c" && i+1 < len(os.Args)-1 {
				cat = os.Args[i+2]
				break
			}
			if strings.HasPrefix(arg, "-c=") {
				cat = strings.TrimPrefix(arg, "-c=")
				break
			}
		}
		if cat != "" {
			printCategoryHelp(cat)
		} else {
			printGeneralHelp()
		}
	}

	flag.Parse()

	if precision < 1 {
		precision = 1
	}

	if list {
		printList(category)
		return
	}

	args := flag.Args()
	if len(args) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid value %q\n", args[0])
		os.Exit(1)
	}
	from := strings.ToLower(args[1])
	to := strings.ToLower(args[2])

	var result float64
	if category == "temperature" {
		result, err = convertTemperature(value, from, to)
	} else {
		units, ok := unitCategories[category]
		if !ok {
			fmt.Fprintf(os.Stderr, "unknown category %q\n", category)
			os.Exit(1)
		}
		result, err = convert(value, from, to, units)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if verbose {
		fmt.Printf("%s %s = %s %s\n", formatValue(value, precision), from, formatValue(result, precision), to)
	} else {
		fmt.Printf("%s %s\n", formatValue(result, precision), to)
	}
}

// convert converts value from one unit to another using multiplicative factors relative to a common base.
func convert(value float64, from, to string, units map[string]float64) (float64, error) {
	fromFactor, ok := units[from]
	if !ok {
		return 0, fmt.Errorf("unknown unit %q, try --list to see available units", from)
	}
	toFactor, ok := units[to]
	if !ok {
		return 0, fmt.Errorf("unknown unit %q, try --list to see available units", to)
	}
	return value * fromFactor / toFactor, nil
}

// formatValue formats a float64 with the given number of significant digits.
// Uses ×10^n scientific notation for very large or very small values.
// %g may also switch to scientific notation for mid-range values depending on
// precision; those are normalised to ×10^n as well for consistent output.
func formatValue(f float64, precision int) string {
	if f == 0 {
		return "0"
	}
	abs := math.Abs(f)
	if abs >= 1e6 || abs < 1e-4 {
		exp := int(math.Floor(math.Log10(abs)))
		mantissa := f / math.Pow(10, float64(exp))
		mantissaStr := fmt.Sprintf("%.*g", precision, mantissa)
		return fmt.Sprintf("%s\u00d710^%d", mantissaStr, exp)
	}
	s := fmt.Sprintf("%.*g", precision, f)
	// %g uses its own scientific notation when exponent >= precision; convert it.
	if i := strings.IndexByte(s, 'e'); i >= 0 {
		exp, _ := strconv.Atoi(s[i+1:])
		return fmt.Sprintf("%s\u00d710^%d", s[:i], exp)
	}
	return s
}

func printGeneralHelp() {
	cats := make([]string, 0, len(unitCategories))
	for k := range unitCategories {
		cats = append(cats, k)
	}
	sort.Strings(cats)
	fmt.Fprintln(os.Stderr, `Usage: transmute [flags] <value> <from> <to>

Flags:
  -c <category>  Unit category (default: distance)
  -v             Verbose output: shows <value> <from> = <result> <to>
  -p <n>         Significant digits (default: 4)
  --list         List available units for the selected category

Categories:
  `+strings.Join(cats, ", ")+`

Examples:
  transmute 10 km mi
  transmute -c weight 70 kg lb
  transmute -c temperature 100 c f
  transmute -c distance --list
  transmute -v -p 6 -c energy 1 tnt j`)
}

func printCategoryHelp(cat string) {
	examples := map[string]string{
		"distance":    "  transmute 10 km mi\n  transmute 1 ly au\n  transmute 100 ft m",
		"weight":      "  transmute 70 kg lb\n  transmute 1 t kg\n  transmute 16 oz lb",
		"volume":      "  transmute 1 gal l\n  transmute 500 ml floz\n  transmute 1 m3 l",
		"temperature": "  transmute 100 c f\n  transmute 373.15 k c\n  transmute 32 f c",
		"velocity":    "  transmute 100 kmh mps\n  transmute 1 c kmh\n  transmute 340 mps kmh",
		"area":        "  transmute 1 km2 mi2\n  transmute 100 ha acre\n  transmute 1 m2 ft2",
		"pressure":    "  transmute 1 atm pa\n  transmute 14.7 psi pa\n  transmute 1 bar atm",
		"energy":      "  transmute 1 kwh j\n  transmute 1 kcal j\n  transmute 1 ev j",
		"data":        "  transmute 1 gb mb\n  transmute 1024 mib gib\n  transmute 8 b bit",
		"time":        "  transmute 3600 s h\n  transmute 1 y d\n  transmute 1000 ms s",
	}
	ex := examples[cat]
	if ex == "" {
		ex = "  transmute -c " + cat + " <value> <from> <to>"
	}
	fmt.Fprintf(os.Stderr, "Usage: transmute -c %s <value> <from> <to>\n\nExamples:\n%s\n\nUse --list to see available units.\n", cat, ex)
}

func printList(category string) {
	if category == "temperature" {
		keys := make([]string, 0, len(temperatureScales))
		for k := range temperatureScales {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		fmt.Printf("Available temperature units:\n")
		for _, k := range keys {
			fmt.Printf("  %s\n", k)
		}
		return
	}
	units, ok := unitCategories[category]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown category %q\n", category)
		os.Exit(1)
	}
	hidden := unlisted[category]
	listed := make([]string, 0, len(units))
	for k := range units {
		if _, ok := hidden[k]; ok {
			continue
		}
		listed = append(listed, k)
	}
	sort.Strings(listed)
	fmt.Printf("Available %s units:\n", category)
	for _, u := range listed {
		fmt.Printf("  %s\n", u)
	}
}
