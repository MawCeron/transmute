package main

import (
	"encoding/base64"
	"encoding/json"
)

// unlisted tracks unit names that are excluded from --list output, per category.
var unlisted = map[string]map[string]struct{}{}

// encoded is base64-encoded JSON of supplemental unit conversion factors (to SI base).
var encoded = "eyJhcmVhIjp7ImZvb3RiYWxsZmllbGQiOjUzNTEuMCwibWFuaGF0dGFuIjo1OTEwMDAwMC4wLCJ0ZXhhcyI6Njk2MjQxMDAwMDAwLjAsIndhbGVzIjoyMDc3OTAwMDAwMC4wfSwiZGF0YSI6eyJibHVyYXkiOjIwMDAwMDAwMDAwMC4wLCJkdmQiOjM3NjAwMDAwMDAwLjAsImZsb3BweSI6MTE3OTY0ODAuMH0sImRpc3RhbmNlIjp7ImJhbGRlYWdsZSI6Mi4xMzM2LCJiYW5hbmEiOjAuMTc3OCwiYnVzIjoxMy43MTYsImRlYXRoc3RhciI6MTYwMDAwLjAsImVpZmZlbHRvd2VyIjozMzAuMCwiZm9vdGJhbGwiOjkxLjQ0LCJob3Rkb2ciOjAuMTUyNCwid2FzaGluZ21hY2hpbmUiOjAuODV9LCJlbmVyZ3kiOnsibGlnaHRuaW5nIjoxMDAwMDAwMDAwLjAsInNuaWNrZXJzIjoxMTE3MDAwLjAsInRudCI6NDE4NDAwMDAwMC4wfSwicHJlc3N1cmUiOnsiZWxlcGhhbnRfZm9vdCI6NjUwMDAuMCwiZmFydCI6MTIwMDAuMCwic25lZXplIjozMDAwLjB9LCJ0aW1lIjp7ImRvZ2d5ZWFyIjoyMjA3NTIwMDAuMCwibWljcm9jZW50dXJ5IjozMTU1Ljc2LCJuYXAiOjEyMDAuMH0sInZlbG9jaXR5Ijp7ImNvbnRpbmVudGFsZHJpZnQiOjkuNTFlLTEwLCJzbG90aCI6MC4xNSwic25haWwiOjAuMDEzLCJ1c2FpbmJvbHQiOjEyLjR9LCJ2b2x1bWUiOnsiYmF0aHR1YiI6MC4yODQsImNhbiI6MC4wMDAzNTUsImtlZyI6MC4wNTg2NzQsIm9seW1waWNwb29sIjoyNTAwLjAsIndpbmUiOjAuMDAwNzV9LCJ3ZWlnaHQiOnsiYmFsZGVhZ2xlIjo0LjEsImJhbmFuYSI6MC4xMTgsImJsdWV3aGFsZSI6MTUwMDAwLjAsImNhdCI6NC41LCJjb3JnaSI6MTMuNiwiZWxlcGhhbnQiOjUwMDAuMCwiZWlmZmVsIjo3MzAwMDAwLjAsImhvdGRvZyI6MC4wNTcsIm1vb24iOjcuMzQyZSsyMn19"

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
