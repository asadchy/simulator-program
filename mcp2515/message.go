// MCP2515 Stand-Alone CAN Interface
package mcp2515

type Message struct {
	id uint32
	extended bool
	length uint8
	data [8]uint8
}
