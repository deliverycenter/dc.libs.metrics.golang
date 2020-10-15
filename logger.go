package dc_metrics

import (
	"github.com/golang/protobuf/proto"

	dcpubsub "github.com/deliverycenter/dc.libs.metrics.golang/pubsub"

	"github.com/deliverycenter/dc.libs.metrics.golang/protos"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	environment    string
	caller         string
	client         *dcpubsub.PubSub
	metricsDefault Metrics
}

func (l *Logger) Debug(message string, metrics Metrics) {
	l.log("DEBUG", message, metrics)
}

func (l *Logger) Info(message string, metrics Metrics) {
	l.log("INFO", message, metrics)
}

func (l *Logger) Warn(message string, metrics Metrics) {
	l.log("WARN", message, metrics)
}

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

	metricsRequest := &protos.WriteMetricsRequest{
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

	encodedMessage, err := proto.Marshal(metricsRequest)
	if err != nil {
		logrus.WithError(err).Error("failed to marshal metricsRequest as proto message")
	}

	result, err := l.client.Publish(encodedMessage)
	if err != nil {
		logrus.WithError(err).Error("failed to write dc_metrics metrics")
	}
	logrus.Debug(result)
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
