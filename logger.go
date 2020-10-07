package dc_metrics

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Logger struct {
	environment    string
	caller         string
	conn           *grpc.ClientConn
	client         MetricsServiceClient
	metricsDefault Metrics
}

// Debug logs a message at level Debug.
func (l *Logger) Debug(message string, metrics Metrics) {
	l.log("DEBUG", message, metrics)
}

// Info logs a message at level Info.
func (l *Logger) Info(message string, metrics Metrics) {
	l.log("INFO", message, metrics)
}

// Warn logs a message at level Warn.
func (l *Logger) Warn(message string, metrics Metrics) {
	l.log("WARN", message, metrics)
}

// Error logs a message at level Error.
func (l *Logger) Error(message string, metrics Metrics) {
	l.log("ERROR", message, metrics)
}

// Log logs a message to both stdout and metrics API
func (l *Logger) log(level string, message string, metrics Metrics) {
	metrics.Level = level
	metrics.Message = message

	l.logToStdout(metrics)
	l.logToMetricsApi(metrics)
}

func (l *Logger) logToStdout(metrics Metrics) {
	l.setDefaults(&metrics)

	logger := logrus.WithFields(logrus.Fields{
		"CorrelationId":     metrics.correlationId(),
		"Environment":       l.environment,
		"Level":             metrics.Level,
		"Direction":         metrics.Direction,
		"SourceType":        metrics.SourceType,
		"SourceName":        metrics.SourceName,
		"Caller":            l.caller,
		"Action":            metrics.Action,
		"CreateTimestamp":   ptypes.TimestampNow(),
		"DurationMs":        metrics.Duration,
		"RootResourceType":  metrics.RootResourceType,
		"ExtRootResourceId": metrics.ExtRootResourceID,
		"IntRootResourceId": metrics.IntRootResourceID,
		"ChildResourceType": metrics.ChildResourceType,
		"ChildResourceId":   metrics.ChildResourceID,
		"ExtStoreId":        metrics.ExtStoreID,
		"IntStoreId":        metrics.IntStoreID,
		"ErrorCode":         metrics.ErrorCode,
	})
	logger.Info(metrics.Message)
}

func (l *Logger) logToMetricsApi(metrics Metrics) {
	l.setDefaults(&metrics)

	metricsRequest := &WriteMetricsRequest{
		CorrelationId:     metrics.correlationId(),
		Environment:       l.environment,
		Level:             metrics.Level,
		Direction:         metrics.Direction,
		SourceType:        metrics.SourceType,
		SourceName:        metrics.SourceName,
		Caller:            l.caller,
		Action:            metrics.Action,
		CreateTimestamp:   ptypes.TimestampNow(),
		DurationMs:        metrics.Duration,
		RootResourceType:  metrics.RootResourceType,
		ExtRootResourceId: metrics.ExtRootResourceID,
		IntRootResourceId: metrics.IntRootResourceID,
		ChildResourceType: metrics.ChildResourceType,
		ChildResourceId:   metrics.ChildResourceID,
		ExtStoreId:        metrics.ExtStoreID,
		IntStoreId:        metrics.IntStoreID,
		ErrorCode:         metrics.ErrorCode,
	}

	_, err := l.client.WriteMetrics(context.Background(), metricsRequest)
	if err != nil {
		logrus.WithError(err).Error("failed to write dc_metrics metrics")
	}
}

func (l *Logger) setDefaults(metrics *Metrics) {
	if metrics.Level == "" {
		metrics.Level = "INFO"
		if l.metricsDefault.Level != "" {
			metrics.Level = l.metricsDefault.Level
		}
	}

	if metrics.Direction == "" {
		metrics.Direction = l.metricsDefault.Direction
	}

	if metrics.SourceType == "" {
		metrics.SourceType = l.metricsDefault.SourceType
	}

	if metrics.SourceName == "" {
		metrics.SourceName = l.metricsDefault.SourceName
	}

	if metrics.Action == "" {
		metrics.Action = l.metricsDefault.Action
	}

	if metrics.Duration == 0 {
		metrics.Duration = l.metricsDefault.Duration
	}

	if metrics.RootResourceType == "" {
		metrics.RootResourceType = l.metricsDefault.RootResourceType
	}

	if metrics.ExtRootResourceID == "" {
		metrics.ExtRootResourceID = l.metricsDefault.ExtRootResourceID
	}

	if metrics.IntRootResourceID == "" {
		metrics.IntRootResourceID = l.metricsDefault.IntRootResourceID
	}

	if metrics.ChildResourceType == "" {
		metrics.ChildResourceType = l.metricsDefault.ChildResourceType
	}

	if metrics.ChildResourceID == "" {
		metrics.ChildResourceID = l.metricsDefault.ChildResourceID
	}

	if metrics.ExtStoreID == "" {
		metrics.ExtStoreID = l.metricsDefault.ExtStoreID
	}

	if metrics.IntStoreID == "" {
		metrics.IntStoreID = l.metricsDefault.IntStoreID
	}

	if metrics.ErrorCode == "" {
		metrics.ErrorCode = l.metricsDefault.ErrorCode
	}
}