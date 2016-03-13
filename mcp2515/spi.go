// MCP2515 Stand-Alone CAN Interface
package mcp2515

import (
	"fmt"
	"github.com/golang/glog"
	"time"
)

func (d *MCP2515) writeRegister(register string, data ...uint8) error {
	address, err := registerAddress(register)
	if err != nil {
		return err
	}

	glog.V(2).Infof("mcp2515: writeRegister %v=%v", register, data)

	command := []uint8{commands["WRITE"], address}
	buffer := append(command, data...)

	return d.Bus.TransferAndReceiveData(buffer)
}

func (d *MCP2515) readRegister(register string, length int) ([]uint8, error) {
	address, err := registerAddress(register)
	if err != nil {
		return nil, err
	}

	command := []uint8{commands["READ"], address}
	buffer := make([]uint8, len(command) + length)
	copy(buffer, command)

	err = d.Bus.TransferAndReceiveData(buffer)
	data := buffer[len(command):]

	glog.V(2).Infof("mcp2515: readRegister %v=%v", register, data)

	return data, err
}

func (d *MCP2515) readStatus() (uint8, error) {
	command := []uint8{commands["READ_STATUS"]}
	buffer := make([]uint8, len(command) + 1)
	copy(buffer, command)

	err := d.Bus.TransferAndReceiveData(buffer)

	data := buffer[len(command)]
	glog.V(4).Infof("mcp2515: readStatus=%v", data)

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
	return err == nil && status&bufferFulls != bufferFulls
}

func (d *MCP2515) receiveMessage(rxBuffer uint8) (*Message, error) {
	glog.V(2).Infof("mcp2515: receive")

	commandName := "READ_RX0"
	if rxBuffer == 1 {
		commandName = "READ_RX1"
	}
	command := []uint8{commands[commandName]}

	// 4 bytes for id, 1 byte for length, 8 bytes for data
	buffer := make([]uint8, len(command) + 13)
	copy(buffer, command)

	err := d.Bus.TransferAndReceiveData(buffer)
	if err != nil {
		return nil, err
	}

	data := buffer[len(command):]
	message := Message{
		Id:       0,
		Extended: (data[1] & (1 << bits["IDE"])) != 0,
		Length:   (data[4] & 0xF),
		Time:     time.Now(),
	}

	if message.Extended {
		// 29 bit extended identifier
		message.Id = (uint32(data[0]) << 21) |
			(uint32(data[1]&0xE0) << 13) |
			(uint32(data[2]&0x03) << 16) |
			uint32(data[3])

	} else {
		// standard 11 bit identifier
		message.Id = (uint32(data[0]) << 3) |
			(uint32(data[1]) >> 5)
	}
	copy(message.Data[:], data[5:])

	return &message, nil
}

func (d *MCP2515) transmitMessage(txBuffer uint8, message *Message) error {
	glog.V(2).Infof("mcp2515: transmit")

	commandName := "WRITE_TX0"
	if txBuffer == 1 {
		commandName = "WRITE_TX1"
	} else if txBuffer == 2 {
		commandName = "WRITE_TX2"
	}
	command := []uint8{commands[commandName]}

	// 4 bytes for id, 1 byte for length, 8 bytes for data
	data := make([]uint8, 13)

	if message.Extended {
		data[0] = uint8(message.Id >> 21)
		data[1] = (uint8(message.Id>>13) & 0xe0) |
			(1 << bits["EXIDE"]) |
			(uint8(message.Id>>16) & 0x03)
		data[2] = uint8(message.Id >> 8)
		data[3] = uint8(message.Id)
	} else {
		data[0] = uint8(message.Id >> 3)
		data[1] = uint8(message.Id << 5)
	}
	data[4] = message.Length
	copy(data[5:], message.Data[:])

	buffer := append(command, data...)

	err := d.Bus.TransferAndReceiveData(buffer)
	if err != nil {
		return err
	}

	// Initiate transmission
	command = []uint8{commands["RTS"] | (1<<txBuffer)}
	return d.Bus.TransferAndReceiveData(command)
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
