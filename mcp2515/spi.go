// MCP2515 Stand-Alone CAN Interface
package mcp2515

import (
	"fmt"
	"github.com/golang/glog"
)

func (d *MCP2515) writeRegister(register string, data ...uint8) error {
	address, err := registerAddress(register)
	if err != nil {
		return err
	}

	glog.V(2).Infof("mcp2515: writeRegister %v", register)

	command := []uint8{commands["WRITE"], address}
	buffer := append(command, data...)

	return d.Bus.TransferAndReceiveData(buffer)
}

func (d *MCP2515) readRegister(register string, length uint8) ([]uint8, error) {
	address, err := registerAddress(register)
	if err != nil {
		return nil, err
	}

	glog.V(2).Infof("mcp2515: readRegister %v", register)

	command := []uint8{commands["READ"], address}
	data := make([]uint8, length)
	buffer := append(command, data...)

	err = d.Bus.TransferAndReceiveData(buffer)
	return data, err
}

func (d *MCP2515) readStatus() (uint8, error) {
	glog.V(2).Infof("mcp2515: readStatus")
	command := []uint8{commands["READ_STATUS"]}
	data := uint8(0)
	buffer := append(command, data)
	err := d.Bus.TransferAndReceiveData(buffer)
	glog.V(2).Infof("status=%v", data)
	return data, err
}

func (d *MCP2515) reset() error {
	glog.V(2).Infof("mcp2515: reset")

	buffer := []uint8{commands["RESET"]}

	return d.Bus.TransferAndReceiveData(buffer)
}

func (d *MCP2515) checkFreeBuffer() bool {
	bufferFulls := uint8((1 << statusBits["TX0REQ"]) |
		(1 << statusBits["TX1REQ"]) |
		(1 << statusBits["TX2REQ"]))
	status, err := d.readStatus()
	return err == nil && (status&bufferFulls) != bufferFulls
}

func (d *MCP2515) receiveMessage() (*Message, bool, error) {
	status, err := d.readStatus()
	if err != nil {
		return nil, false, err
	}

	commandName := ""

	if isBitSet(status, statusBits["RX0IF"]) {
		commandName = "READ_RX0"
	} else if isBitSet(status, statusBits["RX1IF"]) {
		commandName = "READ_RX1"
	} else {
		return nil, false, nil
	}

	command := []uint8{commands[commandName]}

	// 4 bytes for id, 1 byte for length, 8 bytes for data
	data := [13]uint8{}
	buffer := append(command, data[:]...)

	err = d.Bus.TransferAndReceiveData(buffer)
	if err != nil {
		return nil, false, err
	}

	message := Message{
		id:       0,
		extended: (data[1] & 0x8) != 0,
		length:   (data[2] & 0xF),
	}

	if message.extended {
		message.id = (uint32(data[0]) << 21) |
			(uint32(data[1]&0xE0) << 13) |
			(uint32(data[2]&0x03) << 16) |
			uint32(data[3])

	} else {
		// standard 11 bit identifier
		message.id = (uint32(data[0]) << 3) |
			(uint32(data[1]) >> 5)
	}

	return &message, true, nil
}

func registerAddress(register string) (uint8, error) {
	address, ok := registers[register]
	if ok {
		return address, nil
	} else {
		return 0, fmt.Errorf("Register doesn't exist %v", register)
	}
}

func isBitSet(data uint8, bit uint8) bool {
	return (data & (1 << bit)) != 0
}
