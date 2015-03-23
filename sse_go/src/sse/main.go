/*
Package sse implements the primary inner workings of the SSE Reporter.

The primary function is Run(), which starts a scheduler after initialization and registration of the
reporter with the mothership.
*/

package sse

import (
	"bytes"
	"collector"
	"encoding/json"
	"helper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"fmt"
)

var (
	mothership_url = "http://mothership.serverstatusmonitoring.com"
	register_uri   = "/register-service"
	collector_uri  = "/collector"
	status_uri     = "/status"

	collect_frequency_in_seconds = 2       // When to collect a snapshot and store in cache
	report_frequency_in_seconds  = 16      // When to report all snapshots in cache
	version                      = "1.0.0" // The version of SSE this is

	hostname  = ""
	ipAddress = ""

	log_file           = "/var/log/sphire-sse.log"
	configuration_file = "/etc/sse/sse.conf"
	configuration      = new(Configuration)
	server             = new(Server)
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

/*
 StatusBody struct is a direct map to the status reply from the mothership
 */
type StatusBody struct {
	Status string `json:"status"`
}

/*
 Snapshot struct is a collection of other structs which are relayed from the different segments
 of the collector package.
*/
type Snapshot struct {
	CPU     *collector.CPU     `json:"cpu"`
	Disks   *collector.Disks   `json:"disks"`
	Memory  *collector.Memory  `json:"memory"`
	Network *collector.Network `json:"network"`
	System  *collector.System  `json:"system"`
	Time    time.Time          `json:"system_time"`
}

/*
 Server struct implements identifying data about the server.
*/
type Server struct {
	IpAddress       string `json:"ip_address"`
	Hostname        string `json:"hostname"`
	OperatingSystem struct {
		// grepped from cat /etc/issue
		Distributor string `json:"distributor_id`

		// cat /proc/version_signature
		VersionSignature string `json:"version_signature"`

		// cat /proc/version
		Version string `json:"version"`
	} `json:"operating_system"`
	Hardware struct {
		// grepped from lscpu
		Architecture string `json:"architecture"`
		CPUOpMode    string `json:"cpu_op_mode"`
		CPUCount     string `json:"cpu_count"`
		CPUFamily    string `json:"cpu_family"`
		CPUModel     string `json:"cpu_model"`
		CPUMhz       string `json:"cpu_mhz"`
	} `json:"hardware"`
}

/*
 Cache struct implements multiple Snapshot structs. This is cleared after it is reported to the mothership.
 Also includes the program Version and AccountId - the latter of which is gleaned from the configuration.
*/
type Cache struct {
	Node             []*Snapshot `json:"node"`
	Server           *Server     `json:"server"`
	AccountId        string      `json:"account_id"`
	Version          string      `json:"version"`
	OrganizationID   string      `json:"organization_id"`
	OrganizationName string      `json:"organization_name"`
	MachineNickname  string      `json:"machine_nickname"`
}

/*
Run Program entry point which initializes, registers and runs the main scheduler of the
program. Also handles initialization of the global logger.
*/
func Run() {
	// Define the global logger
	logger, err := os.OpenFile(log_file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println(helper.Trace("Unable to secure log: "+log_file, "ERROR"))
		os.Exit(1)
	}
	defer logger.Close()
	log.SetOutput(logger)

	log.Println(helper.Trace("**** Starting program ****", "OK"))

	status := checkStatus()
	if status == false {
		log.Println(helper.Trace("Mothership unreachable. Check your internet connection.", "ERROR"))
		os.Exit(1)
	}

	// Perform system initialization
	_, err = Initialize()
	if err != nil {
		log.Println(helper.Trace("Exiting.", "ERROR"))
		os.Exit(1)
	}

	// Perform the system registration
	log.Println(helper.Trace("Performing registration.", "OK"))
	body, err := Register()
	if err != nil {
		log.Println(helper.Trace("Unable to register this machine"+string(body), "ERROR"))
		os.Exit(1)
	}

	ticker := time.NewTicker(time.Duration(collect_frequency_in_seconds) * time.Second)

	var counter int = 0
	var cache Cache = Cache{
		AccountId:        configuration.Identification.AccountID,
		OrganizationID:   configuration.Identification.OrganizationID,
		OrganizationName: configuration.Identification.OrganizationName,
		MachineNickname:  configuration.Identification.MachineNickname,
		Version:          version,
		Server:           server}

	for {
		<-ticker.C
		if counter > 0 && counter%report_frequency_in_seconds == 0 {
			cache.Sender()
			cache.Node = nil // Clear the Node Cache

			counter = 0
		} else {
			var snapshot Snapshot = Snapshot{}
			cache.Node = append(cache.Node, snapshot.Collector()) // fill in the Snapshot struct and add to the cache
			counter++

			ticker = updateTicker()
		}
	}
}

/*
updateTicker updates the ticker in order to know when to run the codeblock next
*/
func updateTicker() *time.Ticker {
	var updatedSeconds int = time.Now().Second() + collect_frequency_in_seconds
	nextTick := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(),
		time.Now().Minute(), updatedSeconds, time.Now().Nanosecond(), time.Local)
	return time.NewTicker(nextTick.Sub(time.Now()))
}

/*
Initialize attempts to gather all the data for correct program initialization. Loads config, etc.
returns bool and error - if ever false, error will be set, otherwise if bool is true, error is nil.
*/
func Initialize() (bool, error) {
	var err error = nil
	var err_hstnm error = nil
	var architecture string
	var cpuOpMode string
	var cpuCount string
	var cpuFamily string
	var cpuModel string
	var cpuMhz string

	// Attempt to get the server IP address
	ipAddress, err = helper.GetServerExternalIPAddress()
	if err != nil {
		log.Println(helper.Trace("Initialization failed, IP Address unattainable.", "ERROR"))
		return false, err
	}

	// Get the hostname
	hostname, err_hstnm = os.Hostname()
	if err_hstnm != nil {
		hostname_bt, err_hstnm_exec := exec.Command("hostname").Output()
		if err_hstnm_exec == nil {
			hostname = string(hostname_bt)
		}
	}

	// Load and parse configuration file
	file, _ := os.Open(configuration_file)
	err = json.NewDecoder(file).Decode(configuration)

	// Get data about the server and store it in the struct
	distributor, err_distributor := exec.Command("cat", "/etc/issue").Output()
	if err_distributor != nil {
		distributor = []byte{}
	}

	versionSignature, err_versig := exec.Command("cat", "/proc/version_signature").Output()
	if err_versig != nil {
		versionSignature = []byte{}
	}

	version, err_ver := exec.Command("cat", "/proc/version").Output()
	if err_ver != nil {
		version = []byte{}
	}

	hardware_out, err_hwd := exec.Command("lscpu").Output()
	hardware := []string{}
	if err_hwd == nil {
		hardware = strings.Split(string(hardware_out), "\n")
	}

	max := len(hardware) - 1
	for index, line := range hardware {
		split_line := strings.Split(line, ":")
		if index < max { // because index out of range if we dont have this
			key := string(strings.TrimSpace(split_line[0]))
			value := string(strings.TrimSpace(split_line[1]))

			switch key {
			case "Architecture":
				architecture = value
			case "CPU op-mode(s)":
				cpuOpMode = value
			case "CPU(s)":
				cpuCount = value
			case "CPU family":
				cpuFamily = value
			case "Model":
				cpuModel = value
			case "CPU MHz":
				cpuMhz = value
			}
		}
	}

	server = &Server{
		IpAddress: ipAddress,
		Hostname:  hostname,
		OperatingSystem: struct {
			Distributor      string `json:"distributor_id`
			VersionSignature string `json:"version_signature"`
			Version          string `json:"version"`
		}{
			Distributor:      string(distributor),
			VersionSignature: string(versionSignature),
			Version:          string(version),
		},
		Hardware: struct {
			Architecture string `json:"architecture"`
			CPUOpMode    string `json:"cpu_op_mode"`
			CPUCount     string `json:"cpu_count"`
			CPUFamily    string `json:"cpu_family"`
			CPUModel     string `json:"cpu_model"`
			CPUMhz       string `json:"cpu_mhz"`
		}{
			Architecture: architecture,
			CPUOpMode:    cpuOpMode,
			CPUCount:     cpuCount,
			CPUFamily:    cpuFamily,
			CPUModel:     cpuModel,
			CPUMhz:       cpuMhz,
		},
	}

	if err != nil {
		log.Println(helper.Trace("Initialization failed - could not load configuration.", "ERROR"))
		return false, err
	}

	log.Println(helper.Trace("Initialization complete.", "OK"))
	return true, err
}

/*
Register performs a registration of this instance with the mothership
*/
func Register() (string, error) {
	log.Println(helper.Trace("Starting registration.", "OK"))
	var jsonStr = []byte(`{}`)

	// local struct
	registrationObject := map[string]interface{}{
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

	jsonStr, _ = json.Marshal(registrationObject)
	req, err := http.NewRequest("POST", mothership_url+register_uri+"/"+version, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "REG")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var status StatusBody
	_ = json.Unmarshal(body, &status)

	if status.Status == "upgrade" {
		fmt.Println("There is a new version available. Please consider upgrading.")
		log.Println(helper.Trace("There is a new version available. Please consider upgrading.", "OK"))
	}

	log.Println(helper.Trace("Registration complete.", "OK"))
	return string(body), nil
}

/*
Collector collects a snapshot of the system at the time of calling and stores it in
Snapshot struct.
*/
func (Snapshot *Snapshot) Collector() *Snapshot {
	Snapshot.Time = time.Now().Local()

	var CPU collector.CPU = collector.CPU{}
	Snapshot.CPU = CPU.Collect()

	var Disks collector.Disks = collector.Disks{}
	Snapshot.Disks = Disks.Collect()

	var Memory collector.Memory = collector.Memory{}
	Snapshot.Memory = Memory.Collect()

	var Network collector.Network = collector.Network{}
	Snapshot.Network = Network.Collect()

	var System collector.System = collector.System{}
	Snapshot.System = System.Collect()

	return Snapshot
}

/*
Sender sends the data in Cache to the mothership, then clears the Cache struct so that it can
accept new data.
*/
func (Cache *Cache) Sender() bool {
	var jsonStr = []byte(`{}`)

	jsonStr, _ = json.Marshal(Cache)
	req, err := http.NewRequest("POST", mothership_url+collector_uri, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "SND")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(helper.Trace("Unable to complete request", "ERROR"))
		return false
	}
	defer resp.Body.Close()

	read_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(helper.Trace("Unable to complete request"+string(read_body), "ERROR"))
		return false
	}

	return true
}

/*
 checkStatus checks the status of the mothership
*/
func checkStatus() bool {
	var status_body StatusBody

	resp, err := http.Get(mothership_url + status_uri)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &status_body)
	if err == nil && status_body.Status == "ok" {
		return true
	} else {
		log.Println(helper.Trace("Unable to complete status request", "ERROR"))
		return false
	}
}
