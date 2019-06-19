package main

import (
	"os"
	"testing"
)

func TestStatus(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Okay", "OKAY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Status(); got != tt.want {
				t.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = append(os.Args, "foo")
			main()
		})
	}
}
