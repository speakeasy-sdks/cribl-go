// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PreviewResponseBody struct {
	Count                     int64         `json:"count"`
	Events                    []interface{} `json:"events"`
	ProcessingTimeMS          int64         `json:"processingTimeMS"`
	UseFormattedVisualization bool          `json:"useFormattedVisualization"`
}

func (o *PreviewResponseBody) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *PreviewResponseBody) GetEvents() []interface{} {
	if o == nil {
		return []interface{}{}
	}
	return o.Events
}

func (o *PreviewResponseBody) GetProcessingTimeMS() int64 {
	if o == nil {
		return 0
	}
	return o.ProcessingTimeMS
}

func (o *PreviewResponseBody) GetUseFormattedVisualization() bool {
	if o == nil {
		return false
	}
	return o.UseFormattedVisualization
}
