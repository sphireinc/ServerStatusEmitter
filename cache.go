package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	httpClient = &http.Client{}
)

// Cache struct implements multiple Snapshot structs.
// This is cleared after it is reported to the mothership.
// Also includes the program Version and AccountId - the
// latter of which is gleaned from the configuration.
type Cache struct {
	Node         []*Snapshot
	Server       *Server
	ID           string
	Version      string
	Key          string
	Organization string
	Group        string
	Entity       string
}

// Sender sends the data in Cache to the mothership,
// then clears the Cache struct so that it can accept
// new data.
func (Cache *Cache) Sender(collectorURL string) bool {
	jsonStr, err := json.Marshal(Cache)
	if err != nil {
		LogError(errors.New("malformed JSON in cache.Sender()"))
	}

	req, err := http.NewRequest("POST", collectorURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		LogError(err)
		return false
	}

	req.Header.Set("X-Sse-Time", time.Now().UTC().String())
	req.Header.Set("X-Sse-Mode", Conf.Mode)
	req.Header.Set("X-Sse-Entity", Conf.Identification.Entity)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		LogError(err)
		return false
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	readBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		LogError(errors.New("unable to complete request " + string(readBody)))
		return false
	}

	return true
}
