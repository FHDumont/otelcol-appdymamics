package appdynamics

import (
	"fmt"

	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

type MetricsConfig struct {
	LogMetricRecords bool   `mapstructure:"logMetricRecords"`
	Url              string `mapstructure:"url"`
	Prefix           string `mapstructure:"prefix"`
}

type Config struct {
	exporterhelper.TimeoutSettings `mapstructure:",squash"`
	Metrics                        MetricsConfig `mapstructure:"metrics"`
}

func (cfg *Config) Validate() error {
	if cfg.Metrics.LogMetricRecords && cfg.Metrics.Url == "" {
		return fmt.Errorf("the AppDynamics Machine Agent host is required")
	}
	return nil
}
