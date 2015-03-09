package main

import (
	"fmt"
)

var (
	mothership_url               = "http://mothership.serverstatusmonitoring.com"
	collect_frequency_in_seconds = 5  // When to collect a snapshot and store in cache
	report_frequency_in_seconds  = 60 // When to report all snapshots in cache
)

type Snapshot struct {
	Time     string
	Id       string
	Hostname string
	Type     string
	CPU      string
	Disks    string
	Memory   string
	Network  string
	System   string
	Version  string
}

type Cache struct {
	Node *Snapshot
}

func main() {

}
