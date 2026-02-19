package config

import "go.scnd.dev/open/polygon"

type Config struct {
	AppName *string `yaml:"appName"`
}

func (r *Config) GetWebListen() []string {
	return []string{"tcp", ":8080"}
}

func (r *Config) GetFrontendUrl() string {
	return "http://localhost:5173"
}

func (r *Config) GetDatabaseDsn() string {
	return "postgres://postgres:homepage1@localhost:5432/homepage1?sslmode=disable"
}

func (r *Config) GetPolygonConfig() *polygon.Config {
	return &polygon.Config{
		AppName:               nil,
		AppVersion:            nil,
		AppNamespace:          nil,
		AppInstanceId:         nil,
		TelemetryUrl:          nil,
		TelemetryOrganization: nil,
	}
}
