// MCP2515 Stand-Alone CAN Interface
package mcp2515

import (
	"os"
	"os/signal"
	"time"
)

type MsgChan chan *Message
type ErrChan chan error

func RunMessageLoop(d *MCP2515, rxChan MsgChan, txChan MsgChan,
	errChan ErrChan) {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer d.reset()

	for {
		status, err := d.readStatus()
		reportError(err, errChan)

		tryReceiveMessage(d, status, rxChan, errChan)
		tryTransmitMessage(d, status, txChan, errChan)

		select {
		case <-c:
			// Program done
			return
		default:
		}
	}
}
func tryReceiveMessage(d *MCP2515, status uint8,
	rxChan MsgChan, errChan ErrChan) {
	var rxBuffer uint8
	switch {
	case status&(1<<bits["RX0IF"]) != 0:
		rxBuffer = 0
	case status&(1<<bits["RX1IF"]) != 0:
		rxBuffer = 1
	default:
		// no message received; carry on
		return
	}

	rxMessage, err := d.receiveMessage(rxBuffer)
	reportError(err, errChan)

	select {
	case rxChan <- rxMessage:
		// Message forwarded for processing
	default:
		// Message channel full; carry on
	}
}

func tryTransmitMessage(d *MCP2515, status uint8,
	txChan MsgChan, errChan ErrChan) {

	var txBuffer uint8
	switch {
	case status&(1<<statusBits["TX0REQ"]) == 0:
		txBuffer = 0
	case status&(1<<statusBits["TX1REQ"]) == 0:
		txBuffer = 1
	case status&(1<<statusBits["TX2REQ"]) == 0:
		txBuffer = 2
	default:
		// No empty message buffers. Retry transmitting later
		return
	}

	select {
	case txMessage := <-txChan:
		err := d.transmitMessage(txBuffer, txMessage)
		reportError(err, errChan)
	default:
		// nothing to send; carry on
	}
}

func reportError(err error, errChan ErrChan) {
	if err == nil {
		return
	}

	select {
	case errChan <- err:
		// Error reported
	default:
		// Error channel full; carry on
	}
	time.Sleep(100 * time.Millisecond)
}
