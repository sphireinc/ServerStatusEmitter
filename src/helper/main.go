package helper

import (
	"errors"
	"net"
	"runtime"
	"strconv"
	"strings"
	"os/exec"
	"regexp"
)


var (
	ipAddressRegex, _ = regexp.Compile(`\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`)
)

func main() {
}

// GetServerExternalIPAddress gets the server external IP address (public
// IP address) and returns it.
// It returns an empty string and an error if it encounters an error.
func GetServerExternalIPAddress() (string, error) {
	cmdLineIp, _ := exec.Command("hostname", "-I").Output()

	if ipAddressRegex.Match(cmdLineIp) {
		return string(cmdLineIp), nil
	} else {
		interfaces, err := net.Interfaces()
		if err != nil {
			return "", err
		}
		for _, an_interface := range interfaces {
			if an_interface.Flags&net.FlagUp == 0 {
				continue // interface down
			}
			if an_interface.Flags&net.FlagLoopback != 0 {
				continue // loopback interface
			}
			addresses, err := an_interface.Addrs()
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
		return "", errors.New("Connection to external network could not be detected.")
	}
}

// Trace allows us to know which file and which function is executing at the moment.
// It returns a string.
func Trace(message string, status string) string {
	var debug bool = false
	var trace string
	if debug {
		pc := make([]uintptr, 10) // at least 1 entry needed
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		trace = string(file) + "<" + strconv.Itoa(line) + "> " + f.Name() + "(): " + strings.ToUpper(status) + " " + message
	} else {
		trace = strings.ToUpper(status) + " " + message
	}
	return trace
}
