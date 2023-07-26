// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ChartConfigAxis struct {
	XAxis *string  `json:"xAxis,omitempty"`
	YAxis []string `json:"yAxis,omitempty"`
}

func (o *ChartConfigAxis) GetXAxis() *string {
	if o == nil {
		return nil
	}
	return o.XAxis
}

func (o *ChartConfigAxis) GetYAxis() []string {
	if o == nil {
		return nil
	}
	return o.YAxis
}

type ChartConfig struct {
	Axis        *ChartConfigAxis `json:"axis,omitempty"`
	Legend      *Legend          `json:"legend,omitempty"`
	Series      []ChartSeries    `json:"series,omitempty"`
	Settings    *Settings        `json:"settings,omitempty"`
	SingleValue SingleValue      `json:"singleValue"`
	XAxis       *Axis            `json:"xAxis,omitempty"`
}

func (o *ChartConfig) GetAxis() *ChartConfigAxis {
	if o == nil {
		return nil
	}
	return o.Axis
}

func (o *ChartConfig) GetLegend() *Legend {
	if o == nil {
		return nil
	}
	return o.Legend
}

func (o *ChartConfig) GetSeries() []ChartSeries {
	if o == nil {
		return nil
	}
	return o.Series
}

func (o *ChartConfig) GetSettings() *Settings {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *ChartConfig) GetSingleValue() SingleValue {
	if o == nil {
		return SingleValue{}
	}
	return o.SingleValue
}

func (o *ChartConfig) GetXAxis() *Axis {
	if o == nil {
		return nil
	}
	return o.XAxis
}