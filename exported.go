package dc_metrics

import (
	dcpubsub "github.com/deliverycenter/dc.libs.metrics.golang/pubsub"
)

var (
	lg *Logger
)

func Setup(googleProjectId string, pubSubTopicName string, environment string, caller string, metrics Metrics) (err error) {
	lg = &Logger{
		environment:    environment,
		caller:         caller,
		client:         dcpubsub.New(googleProjectId, pubSubTopicName),
		metricsDefault: metrics,
	}

	return nil
}

func Disable() {
	lg = &Logger{
		disabled: true,
	}
}

// Debug logs a message at level Debug.
func Debug(message string, metrics Metrics) {
	lg.Debug(message, metrics)
}

// Info logs a message at level Info.
func Info(message string, metrics Metrics) {
	lg.Info(message, metrics)
}

// Warn logs a message at level Warn.
func Warn(message string, metrics Metrics) {
	lg.Warn(message, metrics)
}

// Error logs a message at level Error.
func Error(message string, metrics Metrics) {
	lg.Error(message, metrics)
}
