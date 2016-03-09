// MCP2515 Stand-Alone CAN Interface
package mcp2515

// SPI commands
var commands = map[string]uint8{
	"RESET":       0xC0,
	"READ":        0x03,
	"READ_RX0":    0x90,
	"READ_RX1":    0x94,
	"WRITE":       0x02,
	"WRITE_TX":    0x40,
	"RTS":         0x80,
	"READ_STATUS": 0xA0,
	"RX_STATUS":   0xB0,
	"BIT_MODIFY":  0x05,
}

// SPI register addresses
var registers = map[string]uint8{
	"RXF0SIDH":  0x00,
	"RXF0SIDL":  0x01,
	"RXF0EID8":  0x02,
	"RXF0EID0":  0x03,
	"RXF1SIDH":  0x04,
	"RXF1SIDL":  0x05,
	"RXF1EID8":  0x06,
	"RXF1EID0":  0x07,
	"RXF2SIDH":  0x08,
	"RXF2SIDL":  0x09,
	"RXF2EID8":  0x0A,
	"RXF2EID0":  0x0B,
	"BFPCTRL":   0x0C,
	"TXRTSCTRL": 0x0D,
	"CANSTAT":   0x0E,
	"CANCTRL":   0x0F,

	"RXF3SIDH": 0x10,
	"RXF3SIDL": 0x11,
	"RXF3EID8": 0x12,
	"RXF3EID0": 0x13,
	"RXF4SIDH": 0x14,
	"RXF4SIDL": 0x15,
	"RXF4EID8": 0x16,
	"RXF4EID0": 0x17,
	"RXF5SIDH": 0x18,
	"RXF5SIDL": 0x19,
	"RXF5EID8": 0x1A,
	"RXF5EID0": 0x1B,
	"TEC":      0x1C,
	"REC":      0x1D,

	"RXM0SIDH": 0x20,
	"RXM0SIDL": 0x21,
	"RXM0EID8": 0x22,
	"RXM0EID0": 0x23,
	"RXM1SIDH": 0x24,
	"RXM1SIDL": 0x25,
	"RXM1EID8": 0x26,
	"RXM1EID0": 0x27,
	"CNF3":     0x28,
	"CNF2":     0x29,
	"CNF1":     0x2A,
	"CANINTE":  0x2B,
	"CANINTF":  0x2C,
	"EFLG":     0x2D,

	"TXB0CTRL": 0x30,
	"TXB0SIDH": 0x31,
	"TXB0SIDL": 0x32,
	"TXB0EID8": 0x33,
	"TXB0EID0": 0x34,
	"TXB0DLC":  0x35,
	"TXB0D0":   0x36,
	"TXB0D1":   0x37,
	"TXB0D2":   0x38,
	"TXB0D3":   0x39,
	"TXB0D4":   0x3A,
	"TXB0D5":   0x3B,
	"TXB0D6":   0x3C,
	"TXB0D7":   0x3D,

	"TXB1CTRL": 0x40,
	"TXB1SIDH": 0x41,
	"TXB1SIDL": 0x42,
	"TXB1EID8": 0x43,
	"TXB1EID0": 0x44,
	"TXB1DLC":  0x45,
	"TXB1D0":   0x46,
	"TXB1D1":   0x47,
	"TXB1D2":   0x48,
	"TXB1D3":   0x49,
	"TXB1D4":   0x4A,
	"TXB1D5":   0x4B,
	"TXB1D6":   0x4C,
	"TXB1D7":   0x4D,

	"TXB2CTRL": 0x50,
	"TXB2SIDH": 0x51,
	"TXB2SIDL": 0x52,
	"TXB2EID8": 0x53,
	"TXB2EID0": 0x54,
	"TXB2DLC":  0x55,
	"TXB2D0":   0x56,
	"TXB2D1":   0x57,
	"TXB2D2":   0x58,
	"TXB2D3":   0x59,
	"TXB2D4":   0x5A,
	"TXB2D5":   0x5B,
	"TXB2D6":   0x5C,
	"TXB2D7":   0x5D,

	"RXB0CTRL": 0x60,
	"RXB0SIDH": 0x61,
	"RXB0SIDL": 0x62,
	"RXB0EID8": 0x63,
	"RXB0EID0": 0x64,
	"RXB0DLC":  0x65,
	"RXB0D0":   0x66,
	"RXB0D1":   0x67,
	"RXB0D2":   0x68,
	"RXB0D3":   0x69,
	"RXB0D4":   0x6A,
	"RXB0D5":   0x6B,
	"RXB0D6":   0x6C,
	"RXB0D7":   0x6D,

	"RXB1CTRL": 0x70,
	"RXB1SIDH": 0x71,
	"RXB1SIDL": 0x72,
	"RXB1EID8": 0x73,
	"RXB1EID0": 0x74,
	"RXB1DLC":  0x75,
	"RXB1D0":   0x76,
	"RXB1D1":   0x77,
	"RXB1D2":   0x78,
	"RXB1D3":   0x79,
	"RXB1D4":   0x7A,
	"RXB1D5":   0x7B,
	"RXB1D6":   0x7C,
	"RXB1D7":   0x7D,
}

// Bit positions for registers
var bits = map[string]uint8{
	// BFPCTRL
	"B1BFS": 5,
	"B0BFS": 4,
	"B1BFE": 3,
	"B0BFE": 2,
	"B1BFM": 1,
	"B0BFM": 0,

	// TXRTSCTRL
	"B2RTS":  5,
	"B1RTS":  4,
	"B0RTS":  3,
	"B2RTSM": 2,
	"B1RTSM": 1,
	"B0RTSM": 0,

	// CANSTAT
	"OPMOD2": 7,
	"OPMOD1": 6,
	"OPMOD0": 5,
	"ICOD2":  3,
	"ICOD1":  2,
	"ICOD0":  1,

	// CANCTRL
	"REQOP2":  7,
	"REQOP1":  6,
	"REQOP0":  5,
	"ABAT":    4,
	"CLKEN":   2,
	"CLKPRE1": 1,
	"CLKPRE0": 0,

	// CNF3
	"WAKFIL":  6,
	"PHSEG22": 2,
	"PHSEG21": 1,
	"PHSEG20": 0,

	// CNF2
	"BTLMODE": 7,
	"SAM":     6,
	"PHSEG12": 5,
	"PHSEG11": 4,
	"PHSEG10": 3,
	"PHSEG2":  2,
	"PHSEG1":  1,
	"PHSEG0":  0,

	// CNF1
	"SJW1": 7,
	"SJW0": 6,
	"BRP5": 5,
	"BRP4": 4,
	"BRP3": 3,
	"BRP2": 2,
	"BRP1": 1,
	"BRP0": 0,

	// CANINTE
	"MERRE": 7,
	"WAKIE": 6,
	"ERRIE": 5,
	"TX2IE": 4,
	"TX1IE": 3,
	"TX0IE": 2,
	"RX1IE": 1,
	"RX0IE": 0,

	// CANINTF
	"MERRF": 7,
	"WAKIF": 6,
	"ERRIF": 5,
	"TX2IF": 4,
	"TX1IF": 3,
	"TX0IF": 2,
	"RX1IF": 1,
	"RX0IF": 0,

	// EFLG
	"RX1OVR": 7,
	"RX0OVR": 6,
	"TXB0":   5,
	"TXEP":   4,
	"RXEP":   3,
	"TXWAR":  2,
	"RXWAR":  1,
	"EWARN":  0,

	// TXBnCTRL (n: 0, 1, 2)
	"ABTF":  6,
	"MLOA":  5,
	"TXERR": 4,
	"TXREQ": 3,
	"TXP1":  1,
	"TXP0":  0,

	// RXB0CTRL
	"RXM1":    6,
	"RXM0":    5,
	"RXRTR":   3,
	"BUKT":    2,
	"BUKT1":   1,
	"FILHIT0": 0,

	// TXBnSIDL (n: 0, 1)
	"EXIDE": 3,

	// RXB1CTRL
	"FILHIT2": 2,
	"FILHIT1": 1,

	// RXBnSIDL (n: 0, 1)
	"SRR": 4,
	"IDE": 3,

	// RXBnDLC (n: 0, 1)
	"RTR":  6,
	"DLC3": 3,
	"DLC2": 2,
	"DLC1": 1,
	"DLC0": 0,
}

var statusBits = map[string]uint8 {
	"RX0IF": 0,
	"RX1IF": 1,
	"TX0REQ": 2,
	"TX0IF": 3,
	"TX1REQ": 4,
	"TX1IF": 5,
	"TX2REQ": 6,
	"TX2IF": 7,
}
