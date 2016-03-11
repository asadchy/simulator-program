// MCP2515 Stand-Alone CAN Interface
package mcp2515

import (
	"time"
)

type Message struct {
	Id       uint32
	Extended bool
	Length   uint8
	Data     [8]uint8
	Time     time.Time
}
