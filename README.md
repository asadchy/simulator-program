# CAN Simulator Program

This is the program I use on the [Raspberry PI CAN simulator][simulator-hw]  to help develop and test the [Carloop open-source car adapter][carloop].

It is written in Go, using [embd, the Embedded Programming Framework in Go][embd].

[The driver for the MCP2515 CAN controller][mcp2515-driver] defines [CAN message structures][can-message-def] and takes care of the [communication with the CAN controller over SPI][mcp2515-spi].

Currently the [main program][main-program] prints any CAN message received and sends a test message every 10 milliseconds.

## Installation

- [Install Go on the Raspberry Pi][install-go]
- Run `go get githbub.com/carloop/simulator-program`
- In the simulator-program directory, compile with `go build`
- With the [simulator board][simulator-hw] attached, run `./simulator-program`

## License

Copyright 2016 Julien Vanier. Distributed under the MIT license. See [LICENSE](/LICENSE) for details.

[carloop]: https://www.carloop.io
[simulator-hw]: https://github.com/carloop/simulator
[embd]: https://github.com/kidoman/embd
[mcp2515-driver]: /mcp2515
[can-message-def]: /mcp2515/message.go
[mcp2515-spi]: /mcp2515/spi.go
[main-program]: /main.go
[install-go]: https://golang.org/doc/install
