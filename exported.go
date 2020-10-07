package dc_metrics

import (
	"github.com/deliverycenter/dc.libs.metrics.golang/protos"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	lg *Logger
)

func Setup(address, environment, caller string, metrics Metrics) (err error) {
	// Set up a connection to the dc_metrics server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return errors.Wrap(err, "couldn't connect to dc_metrics server")
	}

	lg = &Logger{
		environment:    environment,
		caller:         caller,
		conn:           conn,
		client:         protos.NewMetricsServiceClient(conn),
		metricsDefault: metrics,
	}

	return nil
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

// Close closes the connection to the dc_metrics server.
func Close() error {
	return lg.conn.Close()
}
