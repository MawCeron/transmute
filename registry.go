package main

import (
	"encoding/base64"
	"encoding/json"
)

// unlisted tracks unit names that are excluded from --list output, per category.
var unlisted = map[string]map[string]struct{}{}

// encoded is base64-encoded JSON of supplemental unit conversion factors (to SI base).
var encoded = "eyJhcmVhIjp7ImZvb3RiYWxsZmllbGQiOjUzNTEuMCwibWFuaGF0dGFuIjo1OTEwMDAwMC4wLCJ0ZXhhcyI6Njk2MjQxMDAwMDAwLjAsIndhbGVzIjoyMDc3OTAwMDAwMC4wfSwiZGF0YSI6eyJibHVyYXkiOjIwMDAwMDAwMDAwMC4wLCJkdmQiOjM3NjAwMDAwMDAwLjAsImZsb3BweSI6MTE3OTY0ODAuMH0sImRpc3RhbmNlIjp7ImJhbGRlYWdsZSI6Mi4xMzM2LCJiYW5hbmEiOjAuMTc3OCwiYnVzIjoxMy43MTYsImRlYXRoc3RhciI6MTYwMDAwLjAsImVpZmZlbHRvd2VyIjozMzAuMCwiZm9vdGJhbGwiOjkxLjQ0LCJob3Rkb2ciOjAuMTUyNCwid2FzaGluZ21hY2hpbmUiOjAuODV9LCJlbmVyZ3kiOnsibGlnaHRuaW5nIjoxMDAwMDAwMDAwLjAsInNuaWNrZXJzIjoxMTE3MDAwLjAsInRudCI6NDE4NDAwMDAwMC4wfSwicHJlc3N1cmUiOnsiZWxlcGhhbnRfZm9vdCI6NjUwMDAuMCwiZmFydCI6MTIwMDAuMCwic25lZXplIjozMDAwLjB9LCJ0ZW1wZXJhdHVyZSI6eyJwaXp6YSI6MjYwLjAsImJvZHl0ZW1wIjozNy4wLCJzdW5mYWNlIjo1NTA1LjAsInZlbnVzIjo0NjUuMCwibGlnaHRuaW5nIjozMDAwMC4wfSwidGltZSI6eyJkb2dneWVhciI6MjIwNzUyMDAwLjAsIm1pY3JvY2VudHVyeSI6MzE1NS43NiwibmFwIjoxMjAwLjB9LCJ2ZWxvY2l0eSI6eyJjb250aW5lbnRhbGRyaWZ0Ijo5LjUxZS0xMCwic2xvdGgiOjAuMTUsInNuYWlsIjowLjAxMywidXNhaW5ib2x0IjoxMi40fSwidm9sdW1lIjp7ImJhdGh0dWIiOjAuMjg0LCJjYW4iOjAuMDAwMzU1LCJrZWciOjAuMDU4Njc0LCJvbHltcGljcG9vbCI6MjUwMC4wLCJ3aW5lIjowLjAwMDc1fSwid2VpZ2h0Ijp7ImJhbGRlYWdsZSI6NC4xLCJiYW5hbmEiOjAuMTE4LCJibHVld2hhbGUiOjE1MDAwMC4wLCJjYXQiOjQuNSwiY29yZ2kiOjEzLjYsImVsZXBoYW50Ijo1MDAwLjAsImVpZmZlbCI6NzMwMDAwMC4wLCJob3Rkb2ciOjAuMDU3LCJtb29uIjo3LjM0MmUrMjJ9fQ=="

func init() {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return
	}

	var m map[string]map[string]float64
	if err := json.Unmarshal(data, &m); err != nil {
		return
	}

	for cat, units := range m {
		if cat == "temperature" {
			// temperature requires special construction of  tempScale
			if unlisted["temperature"] == nil {
				unlisted["temperature"] = map[string]struct{}{}
			}
			for nombre, celsius := range units {
				c := celsius
				temperatureScales[nombre] = tempScale{
					toK:   func(v float64) float64 { return v*c + 273.15 },
					fromK: func(k float64) float64 { return (k - 273.15) / c },
				}
				unlisted["temperature"][nombre] = struct{}{}
			}
			continue
		}

		// every other categories
		names := map[string]struct{}{}
		for name, factor := range units {
			names[name] = struct{}{}
			if dst, ok := unitCategories[cat]; ok {
				dst[name] = factor
			}
		}
		unlisted[cat] = names
	}
}
