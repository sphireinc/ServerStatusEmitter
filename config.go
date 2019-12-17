package main

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

const (
	uriBase      = "api"
	configFile   = "config.json"
	RegisterURI  = "register"
	CollectorURI = "collector"
	StatusURI    = "status"
	Version      = "1.0"
)

// Config holds our application configuration
type Config struct {
	Mode           string
	Mothership     string
	Log            string
	Identification identification
	Settings       settings
	Reporting      reporting
}

type identification struct {
	ID           string // AccountID
	Key          string // OrganizationID
	Organization string // OrganizationName
	Group        string // Group
	Entity       string // Entity
}

type settings struct {
	Reporting reporting
	System    system
	Disk      disk
}

type disk struct {
	IncludePartitionData bool
}

type system struct {
	Hostname     string
	IPAddress    string
	IncludeUsers bool
}

type reporting struct {
	// CollectFrequencySeconds tells us how often to collect a snapshot and store it in cache
	CollectFrequencySeconds int

	// ReportFrequencySeconds tells us how often to report all snapshots in cache to mothership
	ReportFrequencySeconds int
}

// Load ingests a JSON config file into our Config struct
func (C *Config) Load() {
	jsonFk, err := os.Open(configFile)
	LogFatalError(err)

	defer func() {
		err = jsonFk.Close()
		LogError(err)
	}()

	byteValue, err := ioutil.ReadAll(jsonFk)
	LogFatalError(err)

	err = json.Unmarshal(byteValue, &C)
	LogFatalError(err)
}

// GetURL returns the mothership URL with or without an appended URI
func (C *Config) GetURL(uri string) string {
	if uri != "" {
		u, err := url.Parse(C.Mothership)
		if err != nil {
			LogError(err)
		}
		u.Path = path.Join(u.Path, uriBase)
		u.Path = path.Join(u.Path, Version)
		u.Path = path.Join(u.Path, uri)
		return u.String()
	}
	return C.Mothership
}

// GetRegisterURI returns the Register URL
func (C *Config) GetRegisterURL() string {
	return C.GetURL(RegisterURI)
}

// GetCollectorURL returns the Collector URL
func (C *Config) GetCollectorURL() string {
	return C.GetURL(CollectorURI)
}

// GetStatusURL returns the Status URL
func (C *Config) GetStatusURL() string {
	return C.GetURL(StatusURI)
}

// MarshalJSON returns a JSON representation of our Config struct
func (C *Config) MarshalJSON() ([]byte, error) {
	return json.Marshal(C)
}
