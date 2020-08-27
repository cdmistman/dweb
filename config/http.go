package config

// HTTP is the configurations for the HTTP service.
type HTTP struct {
	ports []Port
}

// Ports returns the list of ports for the HTTP service.
func (http HTTP) Ports() []Port {
	return http.ports
}
