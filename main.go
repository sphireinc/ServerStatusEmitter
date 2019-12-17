package main

import (
	"errors"
	//"fmt"
	"log"
	"os"
	//"os/signal"
	//"time"
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

	//// Perform registration
	//body, err := Register(map[string]interface{}{
	//	"mothership_url":    Conf.Mothership,
	//	"register_url":      Conf.GetRegisterURL(),
	//	"version":           Version,
	//	"collect_frequency": Conf.Settings.Reporting.CollectFrequencySeconds,
	//	"report_frequency":  Conf.Settings.Reporting.ReportFrequencySeconds,
	//	"hostname":          Conf.Settings.System.Hostname,
	//	"ip_address":        Conf.Settings.System.IPAddress
	//}, Conf.GetRegisterURL())
	//
	//LogError(err)
	//
	//// Set up our collector
	//var counter int
	//var snapshot = sse.Snapshot{}
	//var cache = sse.Cache {
	//	AccountID:        Conf.Identification.ID,
	//	OrganizationID:   Conf.Identification.Key,
	//	OrganizationName: Conf.Identification.Organization,
	//	MachineNickname:  Conf.Identification.MachineNickname,
	//	Version:          Version,
	//	Server:           server
	//}
	//
	//type identification struct {
	//	ID           string
	//	Key          string
	//	Organization string
	//	Group        string
	//	Entity       string
	//}
	//
	//ticker := time.NewTicker(time.Duration(Conf.Reporting.CollectFrequencySeconds) * time.Second)
	//death := make(chan os.Signal, 1)
	//signal.Notify(death, os.Interrupt, os.Kill)
	//
	//go func() {
	//	for {
	//		select {
	//		case <-ticker.C: // send the updated time back via to the channel
	//			// reset the snapshot to an empty struct
	//			snapshot = sse.Snapshot{}
	//
	//			// fill in the Snapshot struct and add to the cache
	//			cache.Node = append(cache.Node, snapshot.Collector(Conf.Settings.Disk.IncludePartitionData, Conf.Settings.System.IncludeUsers))
	//			counter++
	//
	//			if counter > 0 && (counter%Conf.Reporting.ReportFrequencySeconds) == 0 {
	//				cache.Sender(Conf.GetCollectorURL())
	//				cache.Node = nil // Clear the Node Cache
	//				counter = 0
	//			}
	//		case <-death:
	//			fmt.Println("died")
	//			return
	//		}
	//	}
	//}()

	return

}
