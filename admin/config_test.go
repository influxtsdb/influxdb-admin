package admin_test

import (
	"testing"

	"github.com/influxtsdb/influxdb-admin/admin"
)

func TestConfig_Parse(t *testing.T) {
	// Parse configuration.
	var c admin.Config
	c.BindAddress = ":8083"
	c.HTTPSEnabled = true
	c.HTTPSCertificate = "/dev/null"

	// Validate configuration.
	if c.BindAddress != ":8083" {
		t.Fatalf("unexpected bind address: %s", c.BindAddress)
	} else if c.HTTPSEnabled != true {
		t.Fatalf("unexpected https enabled: %v", c.HTTPSEnabled)
	} else if c.HTTPSCertificate != "/dev/null" {
		t.Fatalf("unexpected https certificate: %v", c.HTTPSCertificate)
	}
}
