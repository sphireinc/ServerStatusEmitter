/*
Package sse implements the primary inner workings of the SSE Reporter.

The primary function is Run(), which starts a scheduler after initialization and registration of the
reporter with the mothership.
 */

package sse

import (
	"collector"
	"encoding/json"
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
	log_file                     = "/var/log/sphire-sse.log"
	configuration_file           = "/etc/sse/sse.conf"
	configuration                = Configuration{}
)

/*
 Configuration struct is a direct map to the configuration located in the configuration JSON file.
 */
type Configuration struct {
	Identification struct {
		AccountID        string
		OrganizationID   string
		OrganizationName string
		MachineNickname  string
	}
}

/*
 Snapshot struct is a collection of other structs which are relayed from the different segments
 of the collector package.
 */
type Snapshot struct {
	CPU       *collector.CPU
	Disks     *collector.Disks
	Memory    *collector.Memory
	Network   *collector.Network
	System    *collector.System
}

/*
 Cache struct implements multiple Snapshot structs. This is cleared after it is reported to the mothership.
 Also includes the program Version and AccountId - the latter of which is gleaned from the configuration.
 */
type Cache struct {
	Node *Snapshot
	AccountId string
	Version   string
}

// Run Program entry point which initializes, registers and runs the main scheduler of the
// program. Also handles initialization of the global logger.
func Run() {
	// Define the global logger
	logger, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to secure log: " + log_file)
		os.Exit(1)
	}
	defer logger.Close()
	log.SetOutput(logger)

	log.Println()
	log.Println(helper.Trace("**** Starting program ****", "OK"))

	_, err = Initialize()
	if err != nil {
		log.Println(helper.Trace("Exiting.", "ERROR"))
		os.Exit(1)
	}
	Register()
	// sleep(collect_frequency_in_seconds):
	//     collector()
	//     if current_collection == report_frequency_in_seconds:
	//         sender()
}


// Initialize attempts to gather all the data for correct program initialization. Loads config, etc.
// returns bool and error - if ever false, error will be set, otherwise if bool is true, error is nil.
func Initialize() (bool, error) {
	log.Println(helper.Trace("Starting initialization.", "OK"))
	var err error = nil

	// Attempt to get the server IP address
	ipAddress, err = helper.GetServerExternalIPAddress()
	if err != nil {
		log.Println(helper.Trace("Initialization failed, IP Address unattainable.", "ERROR"))
		return false, err
	}

	// Load and parse configuration file
	file, _ := os.Open(configuration_file)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println(helper.Trace("Initialization failed - could not load configuration.", "ERROR"))
		return false, err
	}

	log.Println(helper.Trace("Initialization complete.", "OK"))
	return true, err
}

// Register performs a registration of this instance with the mothership
func Register() {
	log.Println(helper.Trace("Starting registration.", "OK"))

	// TODO: Make call out to /register-service with ipAddress and registration

	log.Println(helper.Trace("Registration complete.", "OK"))
}

// Collector collects a snapshot of the system at the time of calling and stores it in
// Snapshot struct.
func (Snapshot *Snapshot) Collector() *Snapshot {

	return Snapshot
}

// Sender sends the data in Cache to the mothership, then clears the Cache struct so that it can
// accept new data.
func (Cache *Cache) Sender() bool {

	return true
}
