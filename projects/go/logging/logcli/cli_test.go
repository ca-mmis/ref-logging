package logcli

import (
	"reflect"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
		})
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info()
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
			if got := compileTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compileTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExec(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exec(tt.args.args); got != tt.want {
				t.Errorf("Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
