// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LimitsSamples struct {
	MaxSize string `json:"maxSize"`
}

func (o *LimitsSamples) GetMaxSize() string {
	if o == nil {
		return ""
	}
	return o.MaxSize
}

type Limits struct {
	CPUProfileTTL               string                    `json:"cpuProfileTTL"`
	EdgeMetricsCustomExpression *string                   `json:"edgeMetricsCustomExpression,omitempty"`
	EdgeMetricsMode             *EdgeHeartbeatMetricsMode `json:"edgeMetricsMode,omitempty"`
	EnableMetricsPersistence    bool                      `json:"enableMetricsPersistence"`
	EventsMetadataSources       []string                  `json:"eventsMetadataSources,omitempty"`
	ID                          string                    `json:"id"`
	MetricsDirectory            string                    `json:"metricsDirectory"`
	MetricsFieldsBlacklist      []string                  `json:"metricsFieldsBlacklist"`
	MetricsGCPeriod             string                    `json:"metricsGCPeriod"`
	MetricsMaxCardinality       *int64                    `json:"metricsMaxCardinality,omitempty"`
	MetricsMaxDiskSpace         *string                   `json:"metricsMaxDiskSpace,omitempty"`
	MetricsNeverDropList        []string                  `json:"metricsNeverDropList"`
	MetricsWorkerIDBlacklist    []string                  `json:"metricsWorkerIdBlacklist"`
	MinFreeSpace                string                    `json:"minFreeSpace"`
	Samples                     LimitsSamples             `json:"samples"`
}

func (o *Limits) GetCPUProfileTTL() string {
	if o == nil {
		return ""
	}
	return o.CPUProfileTTL
}

func (o *Limits) GetEdgeMetricsCustomExpression() *string {
	if o == nil {
		return nil
	}
	return o.EdgeMetricsCustomExpression
}

func (o *Limits) GetEdgeMetricsMode() *EdgeHeartbeatMetricsMode {
	if o == nil {
		return nil
	}
	return o.EdgeMetricsMode
}

func (o *Limits) GetEnableMetricsPersistence() bool {
	if o == nil {
		return false
	}
	return o.EnableMetricsPersistence
}

func (o *Limits) GetEventsMetadataSources() []string {
	if o == nil {
		return nil
	}
	return o.EventsMetadataSources
}

func (o *Limits) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Limits) GetMetricsDirectory() string {
	if o == nil {
		return ""
	}
	return o.MetricsDirectory
}

func (o *Limits) GetMetricsFieldsBlacklist() []string {
	if o == nil {
		return []string{}
	}
	return o.MetricsFieldsBlacklist
}

func (o *Limits) GetMetricsGCPeriod() string {
	if o == nil {
		return ""
	}
	return o.MetricsGCPeriod
}

func (o *Limits) GetMetricsMaxCardinality() *int64 {
	if o == nil {
		return nil
	}
	return o.MetricsMaxCardinality
}

func (o *Limits) GetMetricsMaxDiskSpace() *string {
	if o == nil {
		return nil
	}
	return o.MetricsMaxDiskSpace
}

func (o *Limits) GetMetricsNeverDropList() []string {
	if o == nil {
		return []string{}
	}
	return o.MetricsNeverDropList
}

func (o *Limits) GetMetricsWorkerIDBlacklist() []string {
	if o == nil {
		return []string{}
	}
	return o.MetricsWorkerIDBlacklist
}

func (o *Limits) GetMinFreeSpace() string {
	if o == nil {
		return ""
	}
	return o.MinFreeSpace
}

func (o *Limits) GetSamples() LimitsSamples {
	if o == nil {
		return LimitsSamples{}
	}
	return o.Samples
}
