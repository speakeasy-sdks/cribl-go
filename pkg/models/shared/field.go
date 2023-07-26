// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Field struct {
	Buckets       []Bucket   `json:"buckets"`
	Count         int64      `json:"count"`
	CountDistinct int64      `json:"countDistinct"`
	CountNull     int64      `json:"countNull"`
	Mean          *int64     `json:"mean,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Size          int64      `json:"size"`
	Stdev         *int64     `json:"stdev,omitempty"`
	TopValues     []TopValue `json:"topValues"`
	Type          FieldType  `json:"type"`
}

func (o *Field) GetBuckets() []Bucket {
	if o == nil {
		return []Bucket{}
	}
	return o.Buckets
}

func (o *Field) GetCount() int64 {
	if o == nil {
		return 0
	}
	return o.Count
}

func (o *Field) GetCountDistinct() int64 {
	if o == nil {
		return 0
	}
	return o.CountDistinct
}

func (o *Field) GetCountNull() int64 {
	if o == nil {
		return 0
	}
	return o.CountNull
}

func (o *Field) GetMean() *int64 {
	if o == nil {
		return nil
	}
	return o.Mean
}

func (o *Field) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Field) GetSize() int64 {
	if o == nil {
		return 0
	}
	return o.Size
}

func (o *Field) GetStdev() *int64 {
	if o == nil {
		return nil
	}
	return o.Stdev
}

func (o *Field) GetTopValues() []TopValue {
	if o == nil {
		return []TopValue{}
	}
	return o.TopValues
}

func (o *Field) GetType() FieldType {
	if o == nil {
		return FieldType("")
	}
	return o.Type
}