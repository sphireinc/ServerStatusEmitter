package main

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

// Register performs a registration of this instance with the mothership
func Register(registrationObject map[string]interface{}, registrationURL string) (string, error) {
	log.Println(helper.Trace(errors.New("starting registration"), "OK"))
	var jsonStr = []byte(`{}`)

	jsonStr, _ = json.Marshal(registrationObject)
	fmt.Println(string(registrationURL))
	req, err := http.NewRequest("POST", registrationURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "REG")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var status helper.Status
	_ = json.Unmarshal(body, &status)

	if status.Status == "upgrade" {
		log.Println(helper.Trace(errors.New("there is a new version available. Please consider upgrading"), "OK"))
		fmt.Println("There is a new version available. Please consider upgrading")
	}

	log.Println(helper.Trace(errors.New("registration complete"), "OK"))
	return string(body), nil
}
