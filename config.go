package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var configFile = "config.json"

type Config struct {
	Mothership     string
	Log            string
	Identification Identification
	Settings       Settings
}

type Identification struct {
	AccountID        string
	OrganizationID   string
	OrganizationName string
	MachineNickname  string
}

type Settings struct {
	Reporting Reporting
	System    System
	Disk      Disk
}

type Disk struct {
	IncludePartitionData bool
}

type System struct {
	IncludeUsers bool
}

type Reporting struct {
	CollectFrequencySeconds int
	ReportFrequencySeconds  int
}

func (C *Config) Load() error {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		HandleErrorFatal(err)
	}

	defer func(){
		_ = jsonFile.Close()
	}()

	byteValue, errr := ioutil.ReadAll(jsonFile)
	if errr != nil {
		HandleErrorFatal(err)
	}

	_ = json.Unmarshal(byteValue, &C)

	return nil
}
