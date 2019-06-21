package logging

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func Metrics() (string, error) {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return execMetrics(sess, nil)
}

func MetricsWithFilter(filter MetricFilter) (string, error) {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return execMetrics(sess, &filter)
}

func MetricsWithCredentials(config SessionConfig) (string, error) {

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(config.SessionRegion),
		Credentials: credentials.NewStaticCredentials(config.SessionAccessKey, config.SessionSecretKey, config.SessionToken),
	}))

	return execMetrics(sess, nil)
}

func execMetrics(sess *session.Session, filter *MetricFilter) (string, error) {
	// Create CloudWatch client
	svc := cloudwatch.New(sess)

	var metricsInput = cloudwatch.ListMetricsInput{}
	if filter != nil {
		metricsInput.MetricName = filter.Name
		metricsInput.Namespace = filter.Namespace
	}
	result, err := svc.ListMetrics(&metricsInput)
	if err != nil {
		fmt.Println("Error", err)
		return "", err
	}

	metrics, err := json.Marshal(result.Metrics)
	if err != nil {
		fmt.Println("Error", err)
		return "", err
	}

	return string(metrics), nil
}
