package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var ipAddressRegex, _ = regexp.Compile(`\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`)

// Status struct is a direct map to the status reply from the mothership
type Status struct {
	Status string `json:"status"`
}

// CheckStatus checks the status of the mothership
func (status_body Status) CheckStatus(uri string) bool {
	resp, err := http.Get(uri)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &status_body)
	if err == nil && status_body.Status == "ok" {
		return true
	}

	log.Println(Trace(errors.New("unable to complete status request"), "ERROR"))
	fmt.Println(errors.New("unable to complete status request"), "ERROR")
	return false
}

// GetServerExternalIPAddress gets the server external IP address (public
// IP address) and returns it.
// It returns an empty string and an error if it encounters an error.
func GetServerExternalIPAddress() (string, error) {
	cmdLineIP, _ := exec.Command("hostname", "-I").Output()

	if !ipAddressRegex.Match(cmdLineIP) {
		interfaces, err := net.Interfaces()
		if err != nil {
			return "", err
		}
		for _, anInterface := range interfaces {
			if anInterface.Flags&net.FlagUp == 0 {
				continue // interface down
			}
			if anInterface.Flags&net.FlagLoopback != 0 {
				continue // loopback interface
			}
			addresses, err := anInterface.Addrs()
			if err != nil {
				return "", err
			}
			for _, address := range addresses {
				var ip net.IP
				switch v := address.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip == nil || ip.IsLoopback() {
					continue
				}
				ip = ip.To4()
				if ip == nil {
					continue // not an ipv4 address
				}
				return ip.String(), nil
			}
		}
		return "", errors.New("connection to external network could not be detected")
	}

	return strings.Replace(strings.Replace(string(cmdLineIP), "\n", "", -1), " ", "", -1), nil
}

