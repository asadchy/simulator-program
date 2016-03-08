package main

import (
	"flag"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"

	"github.com/carloop/simulator/mcp2515"
)

func main() {
	flag.Parse()

	err := embd.InitSPI()
	if err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	const (
		device  = 0
		speed   = 1e5
		bpw     = 8
		delay   = 0
		channel = 0
	)

	spi := embd.NewSPIBus(embd.SPIMode0, device, int(speed), bpw, delay)
	defer spi.Close()

	canDevice := mcp2515.New(spi)
	err = canDevice.Setup(500000)

	if err != nil {
		panic(err)
	}
}
