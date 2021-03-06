package common

import (
	"time"
)

const (
	TimeV1Key         = "time"
	TimeV2Key         = "_time"
	TraceIDKey        = "trace_id"
	ServiceNameKey    = "service_name"
	OperationNameKey  = "operation_name"
	SpanIDKey         = "span_id"
	DurationKey       = "duration"
	FlagsKey          = "flags"
	ProcessTagKeysKey = "process_tag_keys"
	ReferencesKey     = "references"

	DefaultSpanMeasurement     = "span"      // Store spans here
	DefaultSpanMetaMeasurement = "span-meta" // Store sparse span meta data here
	DefaultLogMeasurement      = "log"       // Store logs here

	MeasurementKey = "_measurement"
	FieldKey       = "_field"
	ValueKey       = "_value"

	ReferenceTypeChildOf     = "ChildOf"
	ReferenceTypeFollowsFrom = "FollowsFrom"

	MaxFlushPoints   = 5000
	MaxFlushInterval = time.Second

	MetaCacheInterval = time.Hour
)
