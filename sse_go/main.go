package main

import (
	"collector"
	"fmt"
	"helper"
	"log"
	"os"
)

var (
	mothership_url               = "http://mothership.serverstatusmonitoring.com"
	collect_frequency_in_seconds = 1       // When to collect a snapshot and store in cache
	report_frequency_in_seconds  = 5       // When to report all snapshots in cache
	version                      = "1.0.0" // The version of SSE this is
	hostname                     = ""
	ipAddress                    = ""
	accountId                    = ""
	log_file                     = "/var/log/sphire-sse.log"
)

type Snapshot struct {
	AccountId string
	CPU       *collector.CPU
	Disks     *collector.Disks
	Memory    *collector.Memory
	Network   *collector.Network
	System    *collector.System
	Version   string
}

type Cache struct {
	Node *Snapshot
}

func main() {
	// Define the logger
	logger, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to secure log: " + log_file)
		os.Exit(1)
	}
	defer logger.Close()
	log.SetOutput(logger)

	log.Println()
	log.Println(helper.Trace() + "**** Starting program ****")

	initialize()
	register()
	// sleep(collect_frequency_in_seconds):
	//     collector()
	//     if current_collection == report_frequency_in_seconds:
	//         sender()
}

func initialize() (bool, error) {
	log.Println(helper.Trace() + "Starting initialization.")
	var err error = nil

	ipAddress, err = helper.GetServerExternalIPAddress()
	if err != nil {
		log.Println(helper.Trace() + "Initialization failed, IP Address unattainable.")
		return false, err
	}

	// TODO: Load configuration file from /etc/sse/sse.conf

	log.Println(helper.Trace() + "Initialization complete.")
	return true, err
}

func register() {
	log.Println(helper.Trace() + "Starting registration.")

	// TODO: Make call out to /register-service with ipAddress and registration

	log.Println(helper.Trace() + "Registration complete.")
}

func (Snapshot *Snapshot) collector() *Snapshot {

	return Snapshot
}

func (Cache *Cache) sender() bool {

	return true
}
