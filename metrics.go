package dc_metrics

import (
	"strings"
)

type Metrics struct {
	Message           string
	Level             string
	Direction         string
	SourceType        string
	SourceName        string
	Action            string
	Duration          int32
	RootResourceType  string
	ExtRootResourceID string
	IntRootResourceID string
	ChildResourceType string
	ChildResourceID   string
	ExtStoreID        string
	IntStoreID        string
	ErrorCode         string
}

func (metrics *Metrics) correlationId() string {
	correlationIDFields := []string{
		metrics.SourceType,
		metrics.SourceName,
		metrics.ExtRootResourceID,
		metrics.IntStoreID,
	}

	return strings.ToUpper(strings.Join(correlationIDFields, "-"))
}
