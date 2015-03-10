package main

import (
	"helper"
    "collector"
	"fmt"
)

var (
	mothership_url               = "http://mothership.serverstatusmonitoring.com"
	collect_frequency_in_seconds = 1       // When to collect a snapshot and store in cache
	report_frequency_in_seconds  = 5       // When to report all snapshots in cache
	version                      = "1.0.0" // The version of SSE this is
	hostname                     = ""
	ipAddress                    = ""
	accountId                    = ""
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
	initialize()
	// register()
	// sleep(collect_frequency_in_seconds):
	//     collector()
	//     if current_collection == report_frequency_in_seconds:
	//         sender()
}

func initialize() (bool, error) {
	ipAddress, err := helper.ExternalIP()
	if err != nil {
		fmt.Println(err)
        return false, err
	}
	fmt.Println(ipAddress)
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
