package config

const (
	linuxDefaultConfig string = "/etc/dweb"
)

// Config is the overall configuration settings for dweb.
// From here, all configs for the different parts of dweb are loaded.
type Config struct {
}

// LoadConfig loads the configuration from the specified file.
func LoadConfig(from string) (cfg *Config, err error) {
	return
}
