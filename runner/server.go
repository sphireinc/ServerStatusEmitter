package runner

import (
	"errors"
	"github.com/jsanc623/ServerStatusEmitter/helper"
	error2 "github.com/jsanc623/ServerStatusEmitter/sphlog"
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
// initialization. Loads config, etc. Returns bool and sphlog -
// if ever false, sphlog will be set, otherwise if bool is true, sphlog is nil.
func (server *Server) Initialize() (string, string, error) {
	var architecture string
	var cpuOpMode string
	var cpuCount string
	var cpuFamily string
	var cpuModel string
	var cpuMhz string

	// Attempt to get the server IP address
	ipAddress, err := helper.GetServerExternalIPAddress()
	if err != nil {
		error2.LogWarn("Initialize() could not obtain IP address, setting to localhost")
		ipAddress = "localhost"
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
	var hardware []string
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

	server.IPAddress = ipAddress
	server.Hostname = hostname
	server.OperatingSystem.Distributor = strings.TrimSpace(cleanStrRgx.ReplaceAllString(string(distributor), ""))
	server.OperatingSystem.VersionSignature = strings.TrimSpace(cleanStrRgx.ReplaceAllString(string(versionSignature), ""))
	server.OperatingSystem.Version = strings.TrimSpace(cleanStrRgx.ReplaceAllString(string(version), ""))
	server.Hardware.Architecture = architecture
	server.Hardware.CPUOpMode = cpuOpMode
	server.Hardware.CPUCount = cpuCount
	server.Hardware.CPUFamily = cpuFamily
	server.Hardware.CPUModel = cpuModel
	server.Hardware.CPUMhz = cpuMhz

	if err != nil {
		error2.LogFatalError(errors.New("initialization failed"))
		return ipAddress, string(hostname), err
	}

	error2.LogInfo("Initialize() complete")
	return ipAddress, string(hostname), nil
}
