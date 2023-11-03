package appdynamics

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricsExporter struct {
	logger     *zap.Logger
	config     *Config
	appdClient *http.Client
}

type AppDynamicsCustomMetric struct {
	MetricName      string  `json:"metricName"`
	AggregationType string  `json:"aggregatorType"`
	Value           float64 `json:"value"`
}

func newMetricsExporter(cfg *Config, logger *zap.Logger) (*metricsExporter, error) {
	return &metricsExporter{
		logger: logger,
		config: cfg,
	}, nil
}

func (e *metricsExporter) start(ctx context.Context, host component.Host) error {
	var timeout = time.Second * 10
	if e.config.TimeoutSettings.Timeout.String() != "" {
		timeout = time.Duration(e.config.TimeoutSettings.Timeout.Seconds())
	}

	e.appdClient = &http.Client{Timeout: timeout}

	e.logger.Info("Starting AppDynamics Exporter\n")

	return nil
}

func (e *metricsExporter) shutdown(ctx context.Context) error {
	e.appdClient.CloseIdleConnections()
	e.logger.Info("Shutting AppDynamics Exporter\n")
	return nil
}

func (e *metricsExporter) pushMetricsData(ctx context.Context, metrics pmetric.Metrics) error {

	appDynamicsMetricsSlice := []AppDynamicsCustomMetric{}

	for iResourcerMetrics := 0; iResourcerMetrics < metrics.ResourceMetrics().Len(); iResourcerMetrics++ {
		resourceMetrics := metrics.ResourceMetrics().At(iResourcerMetrics)
		for iScopeMetrics := 0; iScopeMetrics < resourceMetrics.ScopeMetrics().Len(); iScopeMetrics++ {
			scopeMetrics := resourceMetrics.ScopeMetrics().At(iScopeMetrics).Metrics()
			for iMetric := 0; iMetric < scopeMetrics.Len(); iMetric++ {
				metric := scopeMetrics.At(iMetric)
				appDynamicsMetric := &AppDynamicsCustomMetric{}
				switch metric.Type() {
				case pmetric.MetricTypeGauge:
					appDynamicsMetric.MetricName = e.config.Metrics.Prefix + "|" + strings.ReplaceAll(metric.Name(), ".", "|")
					appDynamicsMetric.Value = metric.Gauge().DataPoints().At(0).DoubleValue()
					appDynamicsMetric.AggregationType = "AVERAGE"
				}

				e.logger.Sugar().Debugf("Received: %s", appDynamicsMetric)

				if appDynamicsMetric.MetricName != "" {
					appDynamicsMetricsSlice = append(appDynamicsMetricsSlice, *appDynamicsMetric)
				}
			}
		}
	}

	metricsJson, _ := json.Marshal(appDynamicsMetricsSlice)
	response, error := e.appdClient.Post(e.config.Metrics.Url, "application/json", bytes.NewBuffer(metricsJson))
	if error != nil {
		e.logger.Sugar().Errorf("Error posting metric tree records to machine agent: %v", error)
		return error
	}
	defer response.Body.Close()

	return error
}
