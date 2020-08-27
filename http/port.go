package http

import "github.com/cdmistman/dweb/config"

func handlePort(port *config.Port) (recv <-chan []byte, send chan<- []byte, errors chan error) {
	portToRun := make(chan []byte)
	recv = portToRun

	runToPort := make(chan []byte)
	send = runToPort

	errors = make(chan error)

	go func(received chan<- []byte, sending <-chan []byte, errors chan error) {
	}(portToRun, runToPort, errors)

	return
}
