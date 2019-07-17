package main

import (
	"os"
	"testing"
)

func Test_constructBrokerEndpoint(t *testing.T) {
	tests := []struct {
		name                   string
		brokerEnv              string
		portEnv                string
		expectedBrokerEndpoint string
	}{
		{
			name:                   "Default endpoint expected",
			brokerEnv:              "",
			portEnv:                "",
			expectedBrokerEndpoint: "kafka:9092",
		},
		{
			name:                   "Custom broker host with default port",
			brokerEnv:              "kafka.com",
			portEnv:                "",
			expectedBrokerEndpoint: "kafka.com:9092",
		},
		{
			name:                   "Default broker host with custom port",
			brokerEnv:              "",
			portEnv:                "1234",
			expectedBrokerEndpoint: "kafka:1234",
		},
		{
			name:                   "Custom broker host with custom port",
			brokerEnv:              "kafka.com",
			portEnv:                "1234",
			expectedBrokerEndpoint: "kafka.com:1234",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			clearEnvironmentVariables()
			if test.brokerEnv != "" {
				os.Setenv("broker_host", test.brokerEnv)
			}

			if test.portEnv != "" {
				os.Setenv("custom_port", test.portEnv)
			}

			brokerEndpoint := constructBrokerEndpoint()
			if brokerEndpoint != test.expectedBrokerEndpoint {
				t.Errorf("Expected broker endpoint: %s got: %s",
					test.expectedBrokerEndpoint,
					brokerEndpoint)
			}
			os.Unsetenv("broker_host")
			os.Unsetenv("custom_port")
		})
	}
}

func clearEnvironmentVariables() {
	os.Unsetenv("broker_host")
	os.Unsetenv("custom_port")
}
