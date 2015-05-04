package sse

import (
	"fmt"
	"errors"
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

/*
 Server struct implements identifying data about the server.
*/
type Server struct {
	IpAddress       string `json:"ip_address"`
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

/*
 Initialize attempts to gather all the data for correct program initialization. Loads config, etc.
 returns bool and error - if ever false, error will be set, otherwise if bool is true, error is nil.
*/
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
		log.Println(helper.Trace(errors.New("Initialization failed, IP Address unattainable."), "ERROR"))
		fmt.Println("Initialization failed, IP Address unattainable.", "ERROR")
		return server, "", "", "", err
	}

	// Get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname_byte, err := exec.Command("hostname").Output()
		if err == nil {
			hostname = string(hostname_byte)
		}
	}

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
		log.Println(helper.Trace(errors.New("Initialization failed - could not load configuration."), "ERROR"))
		fmt.Println("Initialization failed - could not load configuration.", "ERROR")
		return server, ipAddress, string(hostname), string(version), err
	}

	log.Println(helper.Trace(errors.New("Initialization complete."), "OK"))
	return server, ipAddress, string(hostname), string(version), nil
}
