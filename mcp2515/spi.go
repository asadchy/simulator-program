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

func (d *MCP2515) reset() error {
	glog.V(2).Infof("mcp2515: reset")

	command := []uint8{commands["RESET"]}

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
