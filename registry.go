package main

import (
	"encoding/base64"
	"encoding/json"
)

// unlisted tracks unit names that are excluded from --list output, per category.
var unlisted = map[string]map[string]struct{}{}

// encoded is base64-encoded JSON of supplemental unit conversion factors (to SI base).
var encoded = "eyJhcmVhIjp7ImZvb3RiYWxsZmllbGQiOjUzNTEuMCwibWFuaGF0dGFuIjo1OTEwMDAwMC4wLCJ0ZXhhcyI6Njk2MjQxMDAwMDAwLjAsIndhbGVzIjoyMDc3OTAwMDAwMC4wLCJzb2NjZXJmaWVsZCI6NzE0MC4wLCJ0ZW5uaXNjb3VydCI6MjYwLjAsImJhc2tldGJhbGwiOjQyMC4wLCJjZW50cmFscGFyayI6MzQxMDAwMC4wLCJ2YXRpY2FuIjo0NDAwMDAuMCwiY2hpaHVhaHVhIjoyNDc0NTUwMDAwMDAuMH0sImRhdGEiOnsiYmx1cmF5IjoyMDAwMDAwMDAwMDAuMCwiZHZkIjozNzYwMDAwMDAwMC4wLCJmbG9wcHkiOjExNzk2NDgwLjAsInR3ZWV0IjoyMjQwLjAsImVtYWlsIjo2MDAwMDAuMCwibGludXhpc28iOjIwMDAwMDAwMDAuMCwibmV0ZmxpeDRrIjo3MDAwMDAwMDAwLjAsInNtcyI6MTEyMC4wLCJtcDMiOjIwNDgwMDAwMC4wfSwiZGlzdGFuY2UiOnsiYmFsZGVhZ2xlIjoyLjEzMzYsImJhbmFuYSI6MC4xNzc4LCJidXMiOjEzLjcxNiwiZGVhdGhzdGFyIjoxNjAwMDAuMCwiZWlmZmVsdG93ZXIiOjMzMC4wLCJmb290YmFsbCI6OTEuNDQsImhvdGRvZyI6MC4xNTI0LCJ3YXNoaW5nbWFjaGluZSI6MC44NSwic2Nob29sYnVzIjoxMC4wLCJzb2NjZXJmaWVsZCI6MTA1LjAsImh1bWFuIjoxLjcsImVtcGlyZXN0YXRlIjo0NDMuMCwiYm9laW5nNzQ3Ijo3MC42LCJ0aXRhbmljIjoyNjkuMCwiZ3JlYXR3YWxsIjoyMTE5NjAwMC4wLCJhbWF6b24iOjY0MDAwMDAuMCwibWFyYXRob24iOjQyMTk1LjAsImV2ZXJlc3QiOjg4NDguMCwiYXRhdCI6MjIuNX0sImVuZXJneSI6eyJsaWdodG5pbmciOjEwMDAwMDAwMDAuMCwic25pY2tlcnMiOjExMTcwMDAuMCwidG50Ijo0MTg0MDAwMDAwLjAsImNvZmZlZSI6MTAwMDAwMC4wLCJyZWRidWxsIjo4MDAwMDAuMCwiaXBob25lY2hhcmdlIjo0MDAwMC4wLCJsaWdodGJ1bGIiOjM2MDAwMC4wLCJudWtlIjo0MTg0MDAwMDAwMDAwMDAwLjB9LCJwcmVzc3VyZSI6eyJlbGVwaGFudGZvb3QiOjY1MDAwLjAsImZhcnQiOjEyMDAwLjAsInNuZWV6ZSI6MzAwMC4wfSwidGVtcGVyYXR1cmUiOnsiZnJlZXplciI6LTE4LjAsInBsdXRvIjotMjI5LjAsIm1hcnMiOi02MC4wLCJjb2xkd2MiOi0yMDAuMCwibXlleGhlYXJ0IjotMzAwLjAsInJvb210ZW1wIjoyMi4wLCJib2R5dGVtcCI6MzcuMCwiY29tZm9ydGFibGUiOjIzLjAsImNvZmZlZWhvdCI6NjAuMCwiY2FsZGl0byI6ODUwMC4wLCJvdmVuIjoxODAuMCwicGl6emEiOjI2MC4wLCJhdXRvY2xhdmUiOjEyMS4wLCJ2ZW51cyI6NDY1LjAsImxhdmEiOjEwMDAuMCwic3VuZmFjZSI6NTUwNS4wLCJsaWdodG5pbmciOjMwMDAwLjAsInBpbmNoaWNhbG9yIjoxNTAuMCwiaW5mZXJubyI6NjAwMC4wfSwidGltZSI6eyJkb2dneWVhciI6MjIwNzUyMDAwLjAsIm1pY3JvY2VudHVyeSI6MzE1NS43NiwibmFwIjoxMjAwLjAsImNvZmZlZSI6NjAwLjAsIm1lZXRpbmciOjM2MDAuMCwic3RhbmR1cCI6OTAwLjAsInNwcmludCI6MTIwOTYwMC4wLCJzZW1lc3RlciI6MTU1NTIwMDAuMCwid29ya2RheSI6Mjg4MDAuMCwicHJlZ25hbmN5IjoyMzMyODAwMC4wLCJwYXJzZWMiOjExMDM3NjAwMC4wLCJkb2N0b3IiOjE0MTkxMjAwMC4wfSwidmVsb2NpdHkiOnsiY29udGluZW50YWxkcmlmdCI6OS41MWUtMTAsInNsb3RoIjowLjE1LCJzbmFpbCI6MC4wMTMsInVzYWluYm9sdCI6MTIuNCwiaHVtYW4iOjEuNCwiY2hlZXRhaCI6MjkuMCwiZm9ybXVsYTEiOjgzLjAsImJ1bGxldCI6MzcwLjAsImNvbmNvcmRlIjo2MDMuMCwiaXNzIjo3NjYwLjB9LCJ2b2x1bWUiOnsiYmF0aHR1YiI6MC4yODQsImNhbiI6MC4wMDAzNTUsImtlZyI6MC4wNTg2NzQsIm9seW1waWNwb29sIjoyNTAwLjAsIndpbmUiOjAuMDAwNzUsInNob3QiOjQuNGUtMDUsImNvZmZlZWN1cCI6MC4wMDAyNSwiZnVlbHRhbmsiOjAuMDUsImJsYWRkZXIiOjAuMDAwNSwidGFua2VyIjozMDAuMH0sIndlaWdodCI6eyJiYWxkZWFnbGUiOjQuMSwiYmFuYW5hIjowLjExOCwiYmx1ZXdoYWxlIjoxNTAwMDAuMCwiY2F0Ijo0LjUsImNvcmdpIjoxMy42LCJlbGVwaGFudCI6NTAwMC4wLCJlaWZmZWwiOjczMDAwMDAuMCwiaG90ZG9nIjowLjA1NywibW9vbiI6Ny4zNDJlKzIyLCJodW1hbiI6NzAuMCwiY2FyIjoxNTAwLjAsImVhcnRoIjo1Ljk3MmUrMjQsImNoaWNrZW4iOjIuNSwibmV3Ym9ybiI6My41LCJzdW1vIjoxNDguMCwidHJleCI6ODAwMC4wLCJpc3MiOjQxOTcyNS4wLCJ0YXJkaXMiOjUwMDAwMDAwLjAsImNoaWh1YWh1YSI6Mi43fX0="

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
