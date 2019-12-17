package helper

import (
	"encoding/json"
	"errors"
	"github.com/glendc/go-external-ip"
	"github.com/jsanc623/ServerStatusEmitter/config"
	error2 "github.com/jsanc623/ServerStatusEmitter/error"
	"io/ioutil"
	"net/http"
	"regexp"
)

var Conf config.Config

type Status struct {
	Status string
}

// CheckStatus checks the status of the mothership
func CheckStatus(uri string) error {
	var err error
	var S Status

	resp, err := http.Get(uri)
	if err != nil {
		error2.LogError(err)
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := ioutil.ReadAll(resp.Body)
	jsonUnmarshalErr := json.Unmarshal(body, S)
	if err == nil && jsonUnmarshalErr == nil && S.Status == "ok" {
		return nil
	}

	err = errors.New("unable to complete status request")
	error2.LogError(err)
	return err
}

// GetServerExternalIPAddress gets the server external IP address (public IP address) and returns it.
// If Conf.Settings.System.IPAddress is set, it'll return that instead.
// It returns an empty string and an error if it encounters an error.
func GetServerExternalIPAddress() (string, error) {
	if Conf.Settings.System.IPAddress != "" {
		return Conf.Settings.System.IPAddress, nil
	}

	// Create the default consensus, using the default configuration and no logger.
	consensus := externalip.DefaultConsensus(nil, nil)

	// Get your IP, which is never <nil> when err is <nil>.
	ip, err := consensus.ExternalIP()
	if err != nil {
		return "", errors.New("connection to external network could not be detected")
	}

	return TrimSpaceNewlineInString(ip.String()), nil // return IPv4/IPv6 in string format
}

// TrimSpaceNewlineInString removes newlines
func TrimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(` +\r?\n +`)
	return re.ReplaceAllString(s, " ")
}
