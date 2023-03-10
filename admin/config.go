package admin

const (
	// DefaultBindAddress is the default bind address for the HTTP server.
	DefaultBindAddress = ":8083"
)

// Config represents the configuration for the admin service.
type Config struct {
	BindAddress      string
	HTTPSEnabled     bool
	HTTPSCertificate string
	Version          string
}

// NewConfig returns an instance of Config with defaults.
func NewConfig() Config {
	return Config{
		BindAddress:      DefaultBindAddress,
		HTTPSEnabled:     false,
		HTTPSCertificate: "/etc/ssl/influxdb.pem",
	}
}
