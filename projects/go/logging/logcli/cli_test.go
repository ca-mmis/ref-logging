package logcli

import (
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"Init Case", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			if len(app.Commands) != 1 {
				t.Errorf("app.Commands length = %v, want %v", len(app.Commands), tt.length)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Info Set"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info()
			if !strings.Contains(os.Args[0], app.Name) {
				t.Errorf("app.Name = %v, want %v", app.Name, os.Args[0])
			}
			if app.Usage != Usage {
				t.Errorf("app.Usage = %v, want %v", app.Usage, Usage)
			}
			if app.Author != Author {
				t.Errorf("app.Author = %v, want %v", app.Author, Author)
			}
			if app.Version != Version {
				t.Errorf("app.Version = %v, want %v", app.Version, Version)
			}
			var compileTime = CompileTime()
			if app.Compiled != compileTime {
				t.Errorf("app.CompileTime = %v, want %v", app.Compiled, compileTime)
			}
		})
	}
}

func Test_compileTime(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcCompileTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExec(t *testing.T) {
	type Args struct {
		args []string
	}
	type TestStruct struct {
		name string
		args Args
		want int
	}
	var tests = []TestStruct{}
	var zeroArgs = Args{
		[]string{},
	}
	var execArgOnly = Args{
		[]string{"logcli.test"},
	}
	var versionArg = Args{
		[]string{"logcli.test", "--version"},
	}

	// Had to leave this test out as there is an underlying bug in the CLI project I'm using
	//var invalidArg = Args{
	//	[]string{"logcli.test", "invalid_arg"},
	//}

	tests = append(tests, TestStruct{"Zero Args", zeroArgs, 1})
	tests = append(tests, TestStruct{"Exec Arg Only", execArgOnly, 0})
	tests = append(tests, TestStruct{"Version Arg", versionArg, 0})
	//tests = append(tests, TestStruct{ "Invalid Arg", invalidArg, 3})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exec(tt.args.args); got != tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcCompileTime(t *testing.T) {

	info, err := os.Stat(os.Args[0])
	if err != nil {
		t.Errorf("Test_calcCompileTime error = %v", err.Error())
	}

	tests := []struct {
		name string
		want time.Time
	}{
		{"App Name", info.ModTime()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcCompileTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcCompileTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompileTime(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompileTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompileTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
