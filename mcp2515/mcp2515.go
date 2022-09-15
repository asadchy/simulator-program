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

// Prescalers for a 12 MHz oscillator with a 8 Time Quanta bit time
var prescalers = map[int]int{
	125000: 5,
	250000: 2,
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

	data, err := d.readRegister("CNF2", 1)

	if err != nil {
		return err
	}

	if data[0] != initialCNF2() {
		return errors.New("CAN chip not responding")
	}

	d.writeRegister("RXB0CTRL", initialRXB0CTRL())
	d.writeRegister("CANCTRL", initialCANCTRL())

	d.initialized = true
	return nil
}

// Configures baud rate and synchronization jump width (SJW)
func initialCNF1(prescaler int) uint8 {
	// SJW = 1 TQ
	return uint8(prescaler)
}

// Configure bit timing to have 1 bit = 8 TQ
func initialCNF2() uint8 {
	// PRSEG = 1 TQ
	// PHSEG1 = 3 TQ
	// BTLMODE = 1 => Use CNF3 for PHSEG2
	return (1 << bits["BTLMODE"]) |
		(2 << bits["PHSEG10"]) |
		(0 << bits["PRSEG"])
}

func initialCNF3() uint8 {
	// PHSEG2 = 3 TQ
	return (2 << bits["PHSEG20"])
}

func initialRXB0CTRL() uint8 {
	// Accept all messages
	return (3 << bits["RXM0"]) |
		// Write message into RXB1 if RXB0 is full
		(1 << bits["BUKT"])
}

func initialCANCTRL() uint8 {
	// Start CAN in normal mode
	return 0
}
