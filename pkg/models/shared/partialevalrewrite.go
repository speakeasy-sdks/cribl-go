// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PartialEvalRewrite struct {
	FieldFilter PartialEvalFieldFilter `json:"fieldFilter"`
	NullObj     string                 `json:"nullObj"`
}

func (o *PartialEvalRewrite) GetFieldFilter() PartialEvalFieldFilter {
	if o == nil {
		return PartialEvalFieldFilter{}
	}
	return o.FieldFilter
}

func (o *PartialEvalRewrite) GetNullObj() string {
	if o == nil {
		return ""
	}
	return o.NullObj
}
