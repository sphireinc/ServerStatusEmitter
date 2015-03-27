package main

import (
	"collector"
	"encoding/json"
	"fmt"
	"helper"
	"log"
	"os"
	"sse"
	"time"
)

var (
	mothership_url = "http://mothership.serverstatusmonitoring.com"
	register_uri   = "/register-service"
	collector_uri  = "/collector"
	status_uri     = "/status"

	collect_frequency_in_seconds = 1       // When to collect a snapshot and store in cache
	report_frequency_in_seconds  = 2       // When to report all snapshots in cache
	sse_version                  = "1.0.0" // The version of SSE this is

	hostname  = ""
	ipAddress = ""
	version   = ""

	log_file           = "/var/log/sphire-sse.log"
	configuration_file = "/etc/sse/sse.conf"
	configuration      = new(Configuration)

	CPU     collector.CPU     = collector.CPU{}
	Disks   collector.Disks   = collector.Disks{}
	Memory  collector.Memory  = collector.Memory{}
	Network collector.Network = collector.Network{}
	System  collector.System  = collector.System{}
)

/*
 Configuration struct is a direct map to the configuration located in the configuration JSON file.
*/
type Configuration struct {
	Identification struct {
		AccountID        string `json:"account_id"`
		OrganizationID   string `json:"organization_id"`
		OrganizationName string `json:"organization_name"`
		MachineNickname  string `json:"machine_nickname"`
	} `json:"identification"`
}

func main() {
	// Define the global logger
	logger, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(helper.Trace("Unable to secure log: "+log_file, "ERROR"))
		fmt.Println("Unable to secure log: "+log_file, "ERROR")
		os.Exit(1)
	}
	defer logger.Close()
	log.SetOutput(logger)

	log.Println(helper.Trace("**** Starting program ****", "OK"))

	// Load and parse configuration file
	file, _ := os.Open(configuration_file)
	err = json.NewDecoder(file).Decode(configuration)

	var status helper.Status = helper.Status{}
	var status_result bool = status.CheckStatus(mothership_url + status_uri)
	if status_result == false {
		log.Println(helper.Trace("Mothership unreachable. Check your internet connection.", "ERROR"))
		fmt.Println("Mothership unreachable. Check your internet connection.", "ERROR")
		os.Exit(1)
	}

	// Perform system initialization
	var server_obj sse.Server = sse.Server{}
	var server *sse.Server = &sse.Server{}
	server, ipAddress, hostname, version, err = server_obj.Initialize()
	if err != nil {
		log.Println(helper.Trace("Exiting.", "ERROR"))
		fmt.Println("Exiting.", "ERROR")
		os.Exit(1)
	}

	// Perform the system registration
	log.Println(helper.Trace("Performing registration.", "OK"))

	var registrationObject map[string]interface{} = map[string]interface{}{
		"configuration":     configuration,
		"mothership_url":    mothership_url,
		"register_uri":      register_uri,
		"version":           version,
		"collect_frequency": collect_frequency_in_seconds,
		"report_frequency":  report_frequency_in_seconds,
		"hostname":          hostname,
		"ip_address":        ipAddress,
		"log_file":          log_file,
		"config_file":       configuration_file,
	}
	var registrationUrl string = mothership_url + register_uri + "/" + sse_version
	body, err := sse.Register(registrationObject, registrationUrl)
	if err != nil {
		log.Println(helper.Trace("Unable to register this machine"+string(body), "ERROR"))
		fmt.Println("Unable to register this machine"+string(body), "ERROR")
		os.Exit(1)
	}

	var counter int = 0
	var snapshot sse.Snapshot = sse.Snapshot{}
	var cache sse.Cache = sse.Cache{
		AccountId:        configuration.Identification.AccountID,
		OrganizationID:   configuration.Identification.OrganizationID,
		OrganizationName: configuration.Identification.OrganizationName,
		MachineNickname:  configuration.Identification.MachineNickname,
		Version:          sse_version,
		Server:           server}
	var collectorUrl string = mothership_url + collector_uri

	ticker := time.NewTicker(time.Duration(collect_frequency_in_seconds) * time.Second)

	for {
		<-ticker.C // send the updated time back via the channel

		// reset the snapshot to an empty struct
		snapshot = sse.Snapshot{}

		// fill in the Snapshot struct and add to the cache
		cache.Node = append(cache.Node, snapshot.Collector())
		counter++

		if counter > 0 && counter%report_frequency_in_seconds == 0 {
			cache.Sender(collectorUrl)
			cache.Node = nil // Clear the Node Cache
			counter = 0
		}
	}

	return
}
