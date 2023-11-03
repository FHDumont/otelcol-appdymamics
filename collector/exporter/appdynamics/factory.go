package appdynamics

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	typeStr                = "appdynamics"
	stability              = component.StabilityLevelAlpha
	defaultMachineAngetUrl = "http://localhost:8293/api/v1/metrics"
	defaultPrefix          = "Custom Metrics"
)

func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithMetrics(createMetricsExporter, stability),
		// exporter.WithLogs(createLogsExporter, stability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		Metrics: MetricsConfig{
			LogMetricRecords: false,
			Url:              defaultMachineAngetUrl,
			Prefix:           defaultPrefix,
		},
	}
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {

	oCfg := cfg.(*Config)
	exporter, err := newMetricsExporter(oCfg, set.Logger)
	if err != nil {
		return nil, fmt.Errorf("cannot configure %s metrics exporter: %w", typeStr, err)
	}

	return exporterhelper.NewMetricsExporter(
		ctx,
		set,
		cfg,
		exporter.pushMetricsData,
		exporterhelper.WithStart(exporter.start),
		exporterhelper.WithShutdown(exporter.shutdown),
	)
}
