package main

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	tests := []struct {
		name      string
		inputArgs []string
		wantCmd   command
		wantErr   bool
	}{
		{
			name:      "Valid command with args",
			inputArgs: []string{"gator", "login", "wagslane"},
			wantCmd:   command{Name: "login", Args: []string{"wagslane"}},
			wantErr:   false,
		},
		{
			name:      "Valid command no args",
			inputArgs: []string{"gator", "register"},
			wantCmd:   command{Name: "register", Args: []string{}},
			wantErr:   true,
		},
		{
			name:      "Missing command",
			inputArgs: []string{"gator", "register"},
			wantCmd:   command{Name: "register", Args: []string{}},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCommand(tt.inputArgs)

			// check if we got an error when we expected one
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCommand() error = %v, wantErr %v", got, tt.wantCmd)
				return
			}

			// Check if the commabnd matches what we expect
			if !tt.wantErr && !reflect.DeepEqual(got, tt.wantCmd) {
				t.Errorf("parseCommand() = %v, want %v", got, tt.wantCmd)

			}
		})
	}
}
