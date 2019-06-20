package logging

import (
	"os"
	"testing"
)

func TestMetricsInvalidAWSCredentials(t *testing.T) {
	_, err := Metrics()
	if err == nil {
		t.Error("Expecting error of type InvalidClientTokaenId")
	}
}

func TestMetricsWithCredentials(t *testing.T) {
	session := SessionConfig{
		SessionRegion:    "us-west-2",
		SessionAccessKey: os.Getenv("TEST_AWS_ACCESS_KEY"),
		SessionSecretKey: os.Getenv("TEST_AWS_SECRET_ACCESS_KEY"),
	}
	metrics, err := MetricsWithCredentials(session)
	if err != nil {
		t.Errorf("MetricsWithCredentials() error = %v", err)
		return
	}
	if len(metrics) == 0 {
		t.Error("MetricsWithCredentials() return an invalid result")
	}
}

func TestMetricsWithFilter(t *testing.T) {
	type args struct {
		filter MetricFilter
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MetricsWithFilter(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricsWithFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MetricsWithFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
