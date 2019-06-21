package logging

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"strings"
	"testing"
)

func TestMetricsWithCredentials(t *testing.T) {
	sessionConfig := SessionConfig{
		SessionRegion:    "us-west-2",
		SessionAccessKey: os.Getenv("TEST_AWS_ACCESS_KEY"),
		SessionSecretKey: os.Getenv("TEST_AWS_SECRET_ACCESS_KEY"),
	}
	metrics, err := MetricsWithCredentials(sessionConfig)
	if err != nil {
		t.Errorf("MetricsWithCredentials() error = %v", err)
		return
	}
	if len(metrics) == 0 {
		t.Error("MetricsWithCredentials() return an invalid result")
	}
}

func TestMetricsWithSession(t *testing.T) {
	sessionConfig := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("TEST_AWS_ACCESS_KEY"), os.Getenv("TEST_AWS_SECRET_ACCESS_KEY"), ""),
	}))

	metrics, err := MetricsWithSession(sessionConfig)
	if err != nil {
		t.Errorf("MetricsWithCredentials() error = %v", err)
		return
	}
	if len(metrics) == 0 {
		t.Error("MetricsWithCredentials() return an invalid result")
	}
}

func TestMetricsWithSessionAndFilter(t *testing.T) {
	type Args struct {
		filter MetricFilter
	}
	type TestStruct struct {
		name    string
		args    Args
		want    *string
		wantErr bool
	}
	var tests = []TestStruct{}

	var validFilter = Args{
		MetricFilter{"validFilter", "logging_test"},
	}
	var notItemsFound = Args{
		MetricFilter{"", ""},
	}

	var validFilterResult string
	tests = append(tests, TestStruct{"Valid Filter", validFilter, &validFilterResult, false})
	tests = append(tests, TestStruct{"Invalid Filter", notItemsFound, nil, true})

	sessionConfig := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("TEST_AWS_ACCESS_KEY"), os.Getenv("TEST_AWS_SECRET_ACCESS_KEY"), ""),
	}))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MetricsWithSessionAndFilter(sessionConfig, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("MetricsWithFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if strings.Compare(got, *tt.want) == 0 {
					t.Errorf("MetricsWithFilter() = %v, want %v", got, *tt.want)
				}
			}
		})
	}

	metrics, err := MetricsWithSession(sessionConfig)
	if err != nil {
		t.Errorf("MetricsWithCredentials() error = %v", err)
		return
	}
	if len(metrics) == 0 {
		t.Error("MetricsWithCredentials() return an invalid result")
	}
}
