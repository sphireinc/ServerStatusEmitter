package sse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"helper"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	httpClient = &http.Client{}
)

// Cache struct implements multiple Snapshot structs.
// This is cleared after it is reported to the mothership.
// Also includes the program Version and AccountId - the
// latter of which is gleaned from the configuration.
type Cache struct {
	Node             []*Snapshot `json:"node"`
	Server           *Server     `json:"server"`
	AccountID        string      `json:"account_id"`
	Version          string      `json:"version"`
	OrganizationID   string      `json:"organization_id"`
	OrganizationName string      `json:"organization_name"`
	MachineNickname  string      `json:"machine_nickname"`
}

// Sender sends the data in Cache to the mothership,
// then clears the Cache struct so that it can accept
// new data.
func (Cache *Cache) Sender(collectorURL string) bool {
	var jsonStr = []byte(`{}`)
	jsonStr, _ = json.Marshal(Cache)

	fmt.Println(string(jsonStr))

	req, err := http.NewRequest("POST", collectorURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "SND")
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(helper.Trace(errors.New("Unable to complete request"), "ERROR"))
		return false
	}
	defer resp.Body.Close()

	readBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(helper.Trace(errors.New("unable to complete request "+string(readBody)), "ERROR"))
		fmt.Println("unable to complete request"+string(readBody), "ERROR")
		return false
	}

	return true
}
