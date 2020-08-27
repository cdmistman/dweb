package config

// Port is a port configuration.
type Port struct {
	network, address string
}

// Network returns the network type of the port.
// For more info, see net.Dial.
// Known types are: "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6", "ip", "ip4", and "ip6".
// The other options should probably never be used.
func (port *Port) Network() string {
	return port.network
}

// Address is the address of the port.
// Usually will just be ":{port number}".
func (port *Port) Address() string {
	return port.address
}
