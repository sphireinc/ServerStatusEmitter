package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

// Configuration is the configuration instance (loads the above LogFile)
var Conf = new(Config)

func logger() {
	logger, err := os.OpenFile(Conf.Log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	LogFatalError(err)
	defer func() {
		_ = logger.Close()
	}()
	log.SetOutput(logger)
}

func main() {
	// Load and parse configuration file
	Conf.Load()

	// Define the global logger
	logger()

	var err error
	var server Server

	err = checkStatus(Conf.GetStatusURL())
	if err != nil {
		LogFatalError(errors.New("mothership unreachable - check your configuration"))
	}

	// Perform system initialization
	Conf.Settings.System.IPAddress, Conf.Settings.System.Hostname, err = server.Initialize()
	LogError(err)

	// Perform registration
	_, err = Register(map[string]interface{}{
		"mothership_url":    Conf.Mothership,
		"register_url":      Conf.GetRegisterURL(),
		"version":           Version,
		"collect_frequency": Conf.Settings.Reporting.CollectFrequencySeconds,
		"report_frequency":  Conf.Settings.Reporting.ReportFrequencySeconds,
		"hostname":          Conf.Settings.System.Hostname,
		"ip_address":        Conf.Settings.System.IPAddress,
	}, Conf.GetRegisterURL())

	LogError(err)

	// Set up our collector
	var counter int
	var snapshot Snapshot
	var cache = Cache{
		ID:           Conf.Identification.ID,
		Key:          Conf.Identification.Key,
		Organization: Conf.Identification.Organization,
		Group:        Conf.Identification.Group,
		Entity:       Conf.Identification.Entity,
		Version:      Version,
		Server:       &server,
	}

	ticker := time.NewTicker(time.Duration(Conf.Reporting.CollectFrequencySeconds) * time.Second)
	death := make(chan os.Signal, 1)
	signal.Notify(death, os.Interrupt, os.Kill)

	go func() {
		for {
			select {
			case <-ticker.C: // send the updated time back via to the channel
				// reset the snapshot to an empty struct
				snapshot = Snapshot{}
				snapshot.Collector()

				// fill in the Snapshot struct and add to the cache
				cache.Node = append(cache.Node, &snapshot)
				counter++

				if counter > 0 && (counter%Conf.Reporting.ReportFrequencySeconds) == 0 {
					cache.Sender(Conf.GetCollectorURL())
					cache.Node = nil // Clear the Node Cache
					counter = 0
				}
			case <-death:
				LogInfo("chan died")
				return
			}
		}
	}()

	return

}
