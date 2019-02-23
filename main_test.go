package main

import (
	"errors"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_getKeepaliveConfig(t *testing.T) {
	tests := []struct {
		title                        string
		environmentalKeepaliveConfig string
		expectedConfig               time.Duration
		expectedError                error
	}{
		{
			title: "Keepalive set to 10 minutes",
			environmentalKeepaliveConfig: "10m",
			expectedConfig:               10 * time.Minute,
			expectedError:                nil,
		},
		{
			title: "Keepalive unset, falling back to default 10 seconds",
			environmentalKeepaliveConfig: "",
			expectedConfig:               10 * time.Second,
			expectedError:                nil,
		},
		{
			title: "Keepalive set but not in go duration",
			environmentalKeepaliveConfig: "10",
			expectedConfig:               0,
			expectedError:                errors.New("time: missing unit in duration"),
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			os.Setenv("keepalive_duration", test.environmentalKeepaliveConfig)
			keepaliveDuration, err := getKeepaliveConfig()
			if test.expectedConfig != keepaliveDuration {
				t.Errorf("expected keepalive duration to be: %v got: %v", test.expectedConfig, keepaliveDuration)
			}
			if err != nil {
				if !strings.Contains(err.Error(), test.expectedError.Error()) {
					t.Errorf("expected error from:`%s` type got: %s", test.expectedError.Error(), err.Error())
				}
			}
		})
	}
}
