package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Register performs a registration of this instance with the mothership
func Register(registrationObject map[string]interface{}, registrationURL string) (string, error) {
	LogError(errors.New("starting registration"))
	var jsonStr = []byte(`{}`)

	jsonStr, _ = json.Marshal(registrationObject)

	LogInfo("registration url: " + registrationURL)

	req, err := http.NewRequest("POST", registrationURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		LogError(err)
		LogFatalError(errors.New("could not register this utility with the mothership"))
	}
	req.Header.Set("X-Custom-Header", "REG")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, _ := ioutil.ReadAll(resp.Body)

	var status Status
	_ = json.Unmarshal(body, &status)

	if status.Status == "upgrade" {
		LogError(errors.New("there is a new version available. Please consider upgrading"))
	}

	LogInfo("registration complete")
	return string(body), nil
}
