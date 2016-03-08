// MCP2515 Stand-Alone CAN Interface
package mcp2515

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/kidoman/embd"
)

type MCP2515 struct {
	// Bus to communicate over
	Bus embd.SPIBus

	initialized bool
	baudRate    int
	mu          sync.RWMutex
}

var prescalers = map[int]uint8{
	125000: 7,
	250000: 3,
	500000: 1,
}

// New creates a new MCP2515 CAN driver
func New(bus embd.SPIBus) *MCP2515 {
	return &MCP2515{
		Bus: bus,
	}
}

func (d *MCP2515) Setup(baudRate int) error {
	d.mu.RLock()
	if d.initialized {
		d.mu.RUnlock()
		return nil
	}
	d.mu.RUnlock()

	d.mu.Lock()
	defer d.mu.Unlock()

	glog.V(1).Infof("mcp2515: setup")

	d.baudRate = baudRate
	prescaler, ok := prescalers[baudRate]
	if !ok {
		return fmt.Errorf("Baud rate not supported %v", baudRate)
	}

	// Reset IC
	d.reset()

	// Wait for IC to come back online
	time.Sleep(10 * time.Microsecond)

	// Load CNF3, CNF2, CNF1 in one shot
	d.writeRegister(
		"CNF3",
		initialCNF3(),
		initialCNF2(),
		initialCNF1(prescaler),
	)

	data, err := d.readRegister("CNF1", 1)

	if err != nil {
		return err
	}

	if data[0] != initialCNF1(prescaler) {
		return errors.New("CAN chip not responding")
	}

	d.writeRegister("RXB0CTRL", initialRXB0CTRL())
	d.writeRegister("CANCTRL", initialCANCTRL())

	d.initialized = true
	return nil
}

// Configures baud rate and synchronization jump width (SJW)
func initialCNF1(prescaler uint8) uint8 {
	return prescaler
}

func initialCNF2() uint8 {
	return (1 << bits["BTLMODE"]) |
		(1 << bits["PHSEG11"])
}

func initialCNF3() uint8 {
	return (1 << bits["PHSEG21"])
}

func initialRXB0CTRL() uint8 {
	// Accept all messages
	return (3 << bits["RXM0"])
}

func initialCANCTRL() uint8 {
	// Start CAN in normal mode
	return 0
}
