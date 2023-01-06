package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/influxtsdb/influxdb-admin/admin"
)

var (
	config  admin.Config
	version = "unknown"
	commit  = "unknown"
	build   = "unknown"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
	flag.StringVar(&config.BindAddress, "bind-address", ":8083", "The default bind address used by the admin service")
	flag.BoolVar(&config.HTTPSEnabled, "https-enabled", false, "Whether the admin service should use HTTPS")
	flag.StringVar(&config.HTTPSCertificate, "https-certificate", "", "The SSL certificate used when HTTPS is enabled")
	flag.Parse()
}

func main() {
	log.Printf("InfluxDB admin version: %s, commit: %s, build: %s", version, commit, build)

	config.Version = version
	srv := admin.NewService(config)
	if err := srv.Open(); err != nil {
		log.Printf("Open admin service error: %s\n", err)
		return
	}

	closed := make(chan struct{})
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	log.Print("Listening for signals")

	// Block until one of the signals above is received
	<-signalCh
	log.Print("Signal received, initializing clean shutdown...")
	go func() {
		defer close(closed)
		if srv != nil {
			srv.Close()
		}
	}()

	// Block again until another signal is received, a shutdown timeout elapses,
	// or the Command is gracefully closed
	log.Print("Waiting for clean shutdown...")
	select {
	case <-signalCh:
		log.Print("Second signal received, initializing hard shutdown")
	case <-time.After(time.Second * 30):
		log.Print("Time limit reached, initializing hard shutdown")
	case <-closed:
		log.Print("Server shutdown completed")
	}
}
