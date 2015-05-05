package main

import (
	"collector"
	"encoding/json"
	"errors"
	"fmt"
	"helper"
	"log"
	"os"
	"sse"
	"time"
)

var (
	URL          = "http://mothership.serverstatusmonitoring.com"
	URIRegister  = "/register"
	URICollector = "/collector"
	URIStatus    = "/status"

	Hostname  = ""
	IPAddress = ""

	LogFile           = "/var/log/sphire-sse.log"
	ConfigurationFile = "/etc/sse/sse.conf"
	Configuration     = new(Config)

	CollectFrequencySeconds = 1 // Collect a snapshot and store in cache every X seconds
	ReportFrequencySeconds  = 1 // Report all snapshots in cache every Y seconds

	CPU     collector.CPU     = collector.CPU{}
	Disks   collector.Disks   = collector.Disks{}
	Memory  collector.Memory  = collector.Memory{}
	Network collector.Network = collector.Network{}
	System  collector.System  = collector.System{}

	Version = "1.0.1"
)

/*
 Configuration struct is a direct map to the configuration located in the configuration JSON file.
*/
type Config struct {
	Identification struct {
		AccountID        string `json:"account_id"`
		OrganizationID   string `json:"organization_id"`
		OrganizationName string `json:"organization_name"`
		MachineNickname  string `json:"machine_nickname"`
	} `json:"identification"`
	Settings struct {
		Disk struct {
			IncludePartitionData bool `json:"include_partition_data"`
		} `json:"disk"`
		System struct {
			IncludeUsers bool `json:"include_users"`
		} `json:"system"`
		Reporting struct {
			CollectFrequencySeconds int `json:"collect_frequency_seconds"`
			ReportFrequencySeconds  int `json:"report_frequency_seconds"`
		} `json:"reporting"`
	} `json:"settings"`
}

func main() {
	// Define the global logger
	logger, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	HandleError(err)
	defer logger.Close()
	log.SetOutput(logger)

	// Load and parse configuration file
	file, err := os.Open(ConfigurationFile)
	HandleError(err)
	err = json.NewDecoder(file).Decode(Configuration)
	HandleError(err)

	// Set any parameters that need to be set
	CollectFrequencySeconds = Configuration.Settings.Reporting.CollectFrequencySeconds
	ReportFrequencySeconds = Configuration.Settings.Reporting.ReportFrequencySeconds

	var status helper.Status = helper.Status{}
	var status_result bool = status.CheckStatus(URL + URIStatus)
	if status_result == false {
		HandleError(errors.New("Mothership unreachable. Check your internet connection."))
	}

	// Perform system initialization
	var server_obj sse.Server = sse.Server{}
	server, ipAddress, hostname, version, error := server_obj.Initialize()
	HandleError(error)

	IPAddress = ipAddress
	Hostname = hostname
	Version = version

	// Perform registration
	body, err := sse.Register(map[string]interface{}{
		"configuration":     Configuration,
		"mothership_url":    URL,
		"register_uri":      URIRegister,
		"version":           Version,
		"collect_frequency": CollectFrequencySeconds,
		"report_frequency":  ReportFrequencySeconds,
		"hostname":          Hostname,
		"ip_address":        IPAddress,
		"log_file":          LogFile,
		"config_file":       ConfigurationFile,
	}, URL+URIRegister+"/"+Version)
	if err != nil {
		HandleError(errors.New("Unable to register this machine" + string(body)))
	}

	// Set up our collector
	var counter int = 0
	var snapshot sse.Snapshot = sse.Snapshot{}
	var cache sse.Cache = sse.Cache{
		AccountId:        Configuration.Identification.AccountID,
		OrganizationID:   Configuration.Identification.OrganizationID,
		OrganizationName: Configuration.Identification.OrganizationName,
		MachineNickname:  Configuration.Identification.MachineNickname,
		Version:          Version,
		Server:           server}

	ticker := time.NewTicker(time.Duration(CollectFrequencySeconds) * time.Second)

	for {
		<-ticker.C // send the updated time back via to the channel

		// reset the snapshot to an empty struct
		snapshot = sse.Snapshot{}

		// fill in the Snapshot struct and add to the cache
		cache.Node = append(cache.Node, snapshot.Collector(Configuration.Settings.Disk.IncludePartitionData,
			Configuration.Settings.System.IncludeUsers))
		counter++

		if counter > 0 && counter%ReportFrequencySeconds == 0 {
			cache.Sender(URL + URICollector)
			cache.Node = nil // Clear the Node Cache
			counter = 0
		}
	}

	return

}

func HandleError(err error) {
	if err != nil {
		log.Println(helper.Trace(err, "ERROR"))
		fmt.Println(err, "ERROR")
		os.Exit(1)
	}
}
