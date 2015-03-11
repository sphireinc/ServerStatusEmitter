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
	logger, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to secure log: " + log_file)
		os.Exit(1)
	}
	defer logger.Close()
	log.SetOutput(logger)

	log.Println(helper.Trace() + "Correctly initialized log in main()")

	initialize()
	// register()
	// sleep(collect_frequency_in_seconds):
	//     collector()
	//     if current_collection == report_frequency_in_seconds:
	//         sender()

	fmt.Println(ipAddress)
}

func initialize() (bool, error) {
	var err error = nil

	ipAddress, err = helper.GetServerExternalIPAddress()
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, err
}

func register() {

}

func (Snapshot *Snapshot) collector() *Snapshot {

	return Snapshot
}

func (Cache *Cache) sender() bool {

	return true
}
