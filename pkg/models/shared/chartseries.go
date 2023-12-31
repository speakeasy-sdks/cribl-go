// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ChartSeries struct {
	Color          *string                  `json:"color,omitempty"`
	Data           []map[string]interface{} `json:"data"`
	DataExpression string                   `json:"dataExpression"`
	DataFilter     map[string]interface{}   `json:"dataFilter,omitempty"`
	Name           string                   `json:"name"`
	Type           *ChartType               `json:"type,omitempty"`
}

func (o *ChartSeries) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *ChartSeries) GetData() []map[string]interface{} {
	if o == nil {
		return []map[string]interface{}{}
	}
	return o.Data
}

func (o *ChartSeries) GetDataExpression() string {
	if o == nil {
		return ""
	}
	return o.DataExpression
}

func (o *ChartSeries) GetDataFilter() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.DataFilter
}

func (o *ChartSeries) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ChartSeries) GetType() *ChartType {
	if o == nil {
		return nil
	}
	return o.Type
}
