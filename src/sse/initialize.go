package sse

import (
	"errors"
	"fmt"
	"helper"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	cleanStrRgx, _ = regexp.Compile(`[\n]|(\\+n)|(\\+l)`)
)

// Server struct implements identifying data about the server.
type Server struct {
	IPAddress       string `json:"ip_address"`
	Hostname        string `json:"hostname"`
	OperatingSystem struct {
		// grepped from cat /etc/issue
		Distributor string `json:"distributor_id"`

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

// Initialize attempts to gather all the data for correct program
// initialization. Loads config, etc. Returns bool and error -
// if ever false, error will be set, otherwise if bool is true, error is nil.
func (server *Server) Initialize() (*Server, string, string, string, error) {
	var architecture string
	var cpuOpMode string
	var cpuCount string
	var cpuFamily string
	var cpuModel string
	var cpuMhz string

	// Attempt to get the server IP address
	ipAddress, err := helper.GetServerExternalIPAddress()
	if err != nil {
		log.Println(helper.Trace(errors.New("initialization failed, IP Address unattainable."), "ERROR"))
		fmt.Println("Initialization failed, IP Address unattainable.", "ERROR")
		return server, "", "", "", err
	}

	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostnameByte, err := exec.Command("hostname").Output()
		if err == nil {
			hostname = string(hostnameByte)
		}
	}

	// Get data about the server and store it in the struct
	distributor, errDistributor := exec.Command("cat", "/etc/issue").Output()
	if errDistributor != nil {
		distributor = []byte{}
	}

	versionSignature, errVersig := exec.Command("cat", "/proc/version_signature").Output()
	if errVersig != nil {
		versionSignature = []byte{}
	}

	version, errVer := exec.Command("cat", "/proc/version").Output()
	if errVer != nil {
		version = []byte{}
	}

	hardwareOut, errHwd := exec.Command("lscpu").Output()
	hardware := []string{}
	if errHwd == nil {
		hardware = strings.Split(string(hardwareOut), "\n")
	}

	max := len(hardware) - 1
	for index, line := range hardware {
		splitLine := strings.Split(line, ":")
		if index < max { // because index out of range if we dont have this
			key := string(strings.TrimSpace(splitLine[0]))
			value := string(strings.TrimSpace(splitLine[1]))

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
		IPAddress: ipAddress,
		Hostname:  hostname,
		OperatingSystem: struct {
			Distributor      string `json:"distributor_id"`
			VersionSignature string `json:"version_signature"`
			Version          string `json:"version"`
		}{
			Distributor:      strings.TrimSpace(cleanStrRgx.ReplaceAllString(string(distributor), "")),
			VersionSignature: strings.TrimSpace(cleanStrRgx.ReplaceAllString(string(versionSignature), "")),
			Version:          strings.TrimSpace(cleanStrRgx.ReplaceAllString(string(version), "")),
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
		log.Println(helper.Trace(errors.New("initialization failed - could not load configuration."), "ERROR"))
		fmt.Println("Initialization failed - could not load configuration.", "ERROR")
		return server, ipAddress, string(hostname), string(version), err
	}

	log.Println(helper.Trace(errors.New("initialization complete."), "OK"))
	return server, ipAddress, string(hostname), string(version), nil
}
