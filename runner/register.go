package runner

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/jsanc623/ServerStatusEmitter/helper"
	error2 "github.com/jsanc623/ServerStatusEmitter/sphlog"
	"io/ioutil"
	"net/http"
)

// Register performs a registration of this instance with the mothership
func Register(registrationObject map[string]interface{}, registrationURL string) (string, error) {
	error2.LogError(errors.New("starting registration"))
	var jsonStr = []byte(`{}`)

	jsonStr, _ = json.Marshal(registrationObject)

	error2.LogInfo("registration url: " + registrationURL)

	req, err := http.NewRequest("POST", registrationURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		error2.LogError(err)
		error2.LogFatalError(errors.New("could not register this utility with the mothership"))
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

	var status helper.Status
	_ = json.Unmarshal(body, &status)

	if status.Status == "upgrade" {
		error2.LogError(errors.New("there is a new version available. Please consider upgrading"))
	}

	error2.LogInfo("registration complete")
	return string(body), nil
}
