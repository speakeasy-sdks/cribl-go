// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DistributedSummaryGroups struct {
	Count        int64 `json:"count"`
	Destinations int64 `json:"destinations"`
	Pipelines    int64 `json:"pipelines"`
	Routes       int64 `json:"routes"`
	Sources      int64 `json:"sources"`
}

func (o *DistributedSummaryGroups) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *DistributedSummaryGroups) GetDestinations() int64 {
	if o == nil {
		return 0
	}
	return o.Destinations
}

func (o *DistributedSummaryGroups) GetPipelines() int64 {
	if o == nil {
		return 0
	}
	return o.Pipelines
}

func (o *DistributedSummaryGroups) GetRoutes() int64 {
	if o == nil {
		return 0
	}
	return o.Routes
}

func (o *DistributedSummaryGroups) GetSources() int64 {
	if o == nil {
		return 0
	}
	return o.Sources
}

type DistributedSummaryWorkers struct {
	Alive            int64 `json:"alive"`
	ConfVersions     int64 `json:"confVersions"`
	Count            int64 `json:"count"`
	Groups           int64 `json:"groups"`
	SoftwareVersions int64 `json:"softwareVersions"`
	Unhealthy        int64 `json:"unhealthy"`
}

func (o *DistributedSummaryWorkers) GetAlive() int64 {
	if o == nil {
		return 0
	}
	return o.Alive
}

func (o *DistributedSummaryWorkers) GetConfVersions() int64 {
	if o == nil {
		return 0
	}
	return o.ConfVersions
}

func (o *DistributedSummaryWorkers) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *DistributedSummaryWorkers) GetGroups() int64 {
	if o == nil {
		return 0
	}
	return o.Groups
}

func (o *DistributedSummaryWorkers) GetSoftwareVersions() int64 {
	if o == nil {
		return 0
	}
	return o.SoftwareVersions
}

func (o *DistributedSummaryWorkers) GetUnhealthy() int64 {
	if o == nil {
		return 0
	}
	return o.Unhealthy
}

type DistributedSummary struct {
	Groups  DistributedSummaryGroups  `json:"groups"`
	Workers DistributedSummaryWorkers `json:"workers"`
}

func (o *DistributedSummary) GetGroups() DistributedSummaryGroups {
	if o == nil {
		return DistributedSummaryGroups{}
	}
	return o.Groups
}

func (o *DistributedSummary) GetWorkers() DistributedSummaryWorkers {
	if o == nil {
		return DistributedSummaryWorkers{}
	}
	return o.Workers
}
