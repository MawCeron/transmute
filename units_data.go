package main

var dataUnits = map[string]float64{
	// Bits (SI decimal)
	"bit":  1.0,
	"kbit": 1e3,
	"mbit": 1e6,
	"gbit": 1e9,
	"tbit": 1e12,
	"pbit": 1e15,
	"ebit": 1e18,

	// Bytes (SI decimal) - 1 byte = 8 bits
	"b":  8.0,
	"kb": 8e3,
	"mb": 8e6,
	"gb": 8e9,
	"tb": 8e12,
	"pb": 8e15,
	"eb": 8e18,

	// Bits (IEC binary)
	"kibit": 1024.0,
	"mibit": 1048576.0,
	"gibit": 1073741824.0,
	"tibit": 1099511627776.0,
	"pibit": 1125899906842624.0,

	// Bytes (IEC binary) - kibibyte, mebibyte, etc.
	"kib": 8192.0,
	"mib": 8388608.0,
	"gib": 8589934592.0,
	"tib": 8796093022208.0,
	"pib": 9007199254740992.0,
	"eib": 9.223372036854776e18,

	// Legacy / networking
	"nibble": 4.0,  // half byte
	"word":   16.0, // 2-byte word
	"dword":  32.0, // 4-byte double word
	"qword":  64.0, // 8-byte quad word
}
