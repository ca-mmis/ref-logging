package main

import (
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

func Test_exec(t *testing.T) {
	type Args struct {
		args []string
	}
	type TestStruct struct {
		name string
		args Args
		want int
	}
	var tests = []TestStruct{}
	var execArgOnly = Args{
		[]string{"logcli.test"},
	}
	var versionArg = Args{
		[]string{"logcli.test", "--version"},
	}

	tests = append(tests, TestStruct{"Exec Arg Only", execArgOnly, 0})
	tests = append(tests, TestStruct{"Version Arg", versionArg, 0})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exec(tt.args.args)
			if got != tt.want {
				t.Errorf("exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
