package logging

type SessionConfig struct {
	SessionRegion    string
	SessionAccessKey string
	SessionSecretKey string
	SessionToken     string
}

type MetricFilter struct {
	Name *string
	Namespace *string
}