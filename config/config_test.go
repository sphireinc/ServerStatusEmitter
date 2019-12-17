package config

import (
	"fmt"
	"os"
	"testing"
)

var TestConfig Config

func setUp() {
	TestConfig = Config{
		Mode:       "reporter",
		Mothership: "http://mothership.serverstatusmonitoring.com",
		Log:        "/dev/null",
		Identification: identification{
			ID:           "test-id",
			Key:          "test-key",
			Organization: "test-org",
			Group:        "test-group",
			Entity:       "test-entity",
		},
		Settings: settings{
			Reporting: reporting{
				CollectFrequencySeconds: 1,
				ReportFrequencySeconds:  1,
			},
			System: system{
				Hostname:     "localhost",
				IPAddress:    "127.0.0.1",
				IncludeUsers: false,
			},
			Disk: disk{
				IncludePartitionData: true,
			},
		},
	}
}

func cleanUP() {
	TestConfig = Config{}
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	cleanUP()
	os.Exit(code)
}

func TestConfig_Load(t *testing.T) {

}

func TestConfig_GetURL(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{TestConfig.GetURL(""), "http://mothership.serverstatusmonitoring.com"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if test.actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}

}

func TestConfig_GetStatusURL(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{TestConfig.GetStatusURL(), "http://mothership.serverstatusmonitoring.com/status"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if test.actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}

func TestConfig_GetRegisterURL(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{TestConfig.GetRegisterURL(), "http://mothership.serverstatusmonitoring.com/register"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if test.actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}

func TestConfig_GetCollectorURL(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{TestConfig.GetCollectorURL(), "http://mothership.serverstatusmonitoring.com/collector"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if test.actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}
